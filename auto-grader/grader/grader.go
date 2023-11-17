package grader

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "thing/auto-grader/graderrequest"
)

type Result struct {
	Fileid string

	Username string
	Filename string

	Isgraded      bool
	Obtainedscore *int
	Maxscore      *int
	Feedback      *string
}

type submission struct {
	fileid   string
	username string
	filename string
	isGraded bool
}

type Service interface {
	RegisterFile(username string, filename string) (string, error)
	LaunchGrader(fileid string, filename string)
	LoadResults(username string) ([]Result, error)
	UploadFile(fileid string, filepath string)
}

type service struct {
	conn *pgx.Conn
}

// NewService is used to create a single instance of the service
func NewService(conn *pgx.Conn) Service {
	return &service{
		conn: conn,
	}
}

// In the future, this will make a request to the grader server
func (s *service) LaunchGrader(fileid string, filename string) {
	go s.sendGradeRequest(fileid, filename)
}

func (s *service) UploadFile(fileid string, filepath string) {
	go s.uploadFile(fileid, filepath)
}

func (s *service) RegisterFile(username string, filename string) (string, error) {
	var fileid int
	err := s.conn.QueryRow(
		context.Background(),
		`insert into submissions 
    (username, filename) 
      values ($1, $2) 
    returning fileid;`,
		username, filename).
		Scan(&fileid)
	return strconv.Itoa(fileid), err
}

func (s *service) LoadResults(username string) ([]Result, error) {
	return s.viewGrades(username)
}

func (s *service) viewGrades(username string) ([]Result, error) {
	rows, _ := s.conn.Query(context.Background(),
		`
    select filename, isgraded, obtainedscore, maxscore, feedback 
    from submissions 
    where username=$1
    ;
    `, username)

	var results []Result

	for rows.Next() {
		var result Result

		err := rows.Scan(
			&result.Filename,
			&result.Isgraded,
			&result.Obtainedscore,
			&result.Maxscore,
			&result.Feedback,
		)
		fmt.Printf(result.Username)
		if err != nil {
			return nil, err
		}
		results = append(results, result)

	}
	return results, rows.Err()
}

func (s *service) sendGradeRequest(fileid string, filename string) {
	time.Sleep(time.Second)
	addr := "localhost:4000"

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGraderRequestServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.GradeFile(ctx, &pb.File{Fileid: fileid, Filename: filename})
	if err != nil {
		log.Printf("could not send grade request: %v", err)
	}

	log.Printf("Servercode: %d", r.GetStatusCode())
}

func (s *service) uploadFile(fileid string, filepath string) {
	fmt.Println("Method called")

	bucket := aws.String(os.Getenv("BACKBLAZE_KEY_NAME"))
	key := aws.String(fileid)

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("BACKBLAZE_KEY_ID"),
			os.Getenv("BACKBLAZE_APPLICATION_KEY"),
			"",
		),
		Endpoint:         aws.String(os.Getenv("AWS_ENDPOINT_URL")),
		Region:           aws.String(os.Getenv("AWS_REGION")),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Printf("Error while creaing a new session: %v", err)
		return
	}

	fmt.Println("Session made")
	s3Client := s3.New(newSession)
	fmt.Println("client made")

	filedata, err := os.Open(filepath)
	if err != nil {
		log.Printf("Unable to open file: %s. %v", filepath, err)
		return
	}

	// Upload a new object "testfile.txt" with the string "S3 Compatible API"
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:   filedata,
		Bucket: bucket,
		Key:    key,
	})

	if err != nil {
		fmt.Printf("Failed to upload object %s/%s, %s\n", *bucket, *key, err.Error())
		return
	}
	fmt.Printf("Successfully uploaded key %s\n", *key)
}

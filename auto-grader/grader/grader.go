package grader

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "thing/auto-grader/graderrequest"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jackc/pgx/v5"
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
			// &result.Username,
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

// Bad just for testing
func (s *service) tempTestingGrader(fileid string) {
	// TODO: MAGIC, Makes the grpc call
	time.Sleep(time.Second * 60)

	// TODO: figure out how to remove this variable without causing error
	var filename string

	feedback := "very gud"
	err := s.conn.QueryRow(
		context.Background(),
		`
    UPDATE submissions
    SET isgraded = $1,
        feedback = $2
    WHERE fileid=$3
    returning filename
    ;

    `,
		true, feedback, fileid).Scan(&filename)
	//     RETURNING *;  for getting the object
	if err != nil {
		fmt.Println(err)
	}
}

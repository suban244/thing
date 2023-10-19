package grader

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type Service interface {
	RegisterFile(username string, filename string) (string, error)
	LaunchGrader(fileid string)

	// HandleFileUpload(filename string) error
}

type service struct {
	conn *pgx.Conn
}

type submission struct {
	fileid   string
	username string
	filename string
	isGraded bool
}

// NewService is used to create a single instance of the service
func NewService(conn *pgx.Conn) Service {
	return &service{
		conn: conn,
	}
}

// In the future, this will make a request to the grader server
func (s *service) LaunchGrader(fileid string) {
	go s.grade(fileid)
}

func (s *service) grade(fileid string) {
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

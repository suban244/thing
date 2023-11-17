package db

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type Assignment struct {
	Id          string
	gradingfile string
	createdBy   string

	totalScore *int
}

type Submission struct {
	SubmissionID string

	Username     string
	AssignmentID string

	Isgraded      bool
	Obtainedscore *int
	Feedback      *string
}

type Service interface {
	CreateUser(username string) error
	CheckIfUserExists(username string) (bool, error)
	DeleteUser(username string) error
	UpdateGradingFile(assignmentid, newGradingFileid string) error
	UpdateScore(assignmentid string, score int) error
	GetAllAssignment(username string) ([]Assignment, error)
	ViewSubmissionByUser(username string) ([]Submission, error)
	ViewSubmissionByAssignment(assignmentid string) ([]Submission, error)
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

func (s *service) checkConnection() error {
	var err error
	if s.conn == nil {
		s.conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		return err

	}
	if s.conn.IsClosed() {
		s.conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		return err
	}
	return nil
}

func (s *service) cleanDatabase() error {
	err := s.checkConnection()
	if err != nil {
		return err
	}

	err = s.conn.QueryRow(context.Background(),
		`delete from submission`,
	).Scan()
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}
		fmt.Printf("Could not clear submission table: %v", err)
		return err
	}
	err = s.conn.QueryRow(context.Background(),
		`delete from assignment`,
	).Scan()
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}
		fmt.Printf("Could not clear assignment table: %v", err)
		return err
	}
	err = s.conn.QueryRow(context.Background(),
		`delete from users`,
	).Scan()
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}
		fmt.Printf("Could not clear user table: %v", err)
		return err
	}

	return nil
}

func (s *service) CreateUser(username string) error {
	err := s.checkConnection()
	if err != nil {
		return err
	}

	err = s.conn.QueryRow(
		context.Background(),
		`insert into users (username) values($1);`,
		username,
	).Scan()
	if err == pgx.ErrNoRows {
		return nil
	}
	return err
}

func (s *service) CheckIfUserExists(username string) (bool, error) {
	err := s.checkConnection()
	var i int
	if err != nil {
		return false, err
	}
	err = s.conn.QueryRow(context.Background(),
		`SELECT 1 FROM users WHERE username =  $1;`,
		username).Scan(&i)

	if err == pgx.ErrNoRows {
		return false, nil
	}
	if i == 1 {
		return true, err
	}
	return false, err
}

// func (s *service) getAllUsers() []string

func (s *service) DeleteUser(username string) error {
	err := s.checkConnection()
	if err != nil {
		return err
	}
	err = s.conn.QueryRow(context.Background(),
		`delete from users where username=$1;`,
		username).Scan()
	if err == pgx.ErrNoRows {
		return nil
	}
	return err
}

func (s *service) CreateAssignment(assignmentid, gradingfileid, username string) error {
	err := s.checkConnection()
	if err != nil {
		return err
	}

	err = s.conn.QueryRow(
		context.Background(),
		`insert into assignment (assignmentid, gradingfile, createdby) values ($1, $2, $3);`,
		assignmentid,
		gradingfileid,
		username,
	).Scan()
	if err == pgx.ErrNoRows {
		return nil
	}
	return nil
}

func (s *service) UpdateGradingFile(assignmentid, newGradingFileid string) error {
	err := s.checkConnection()
	if err != nil {
		return err
	}

	err = s.conn.QueryRow(
		context.Background(),
		`update assignment set gradingfile=$1 where assignmentid=$2;`,
		newGradingFileid, assignmentid,
	).Scan()
	if err == pgx.ErrNoRows {
		return nil
	}
	return nil
}

func (s *service) UpdateScore(assignmentid string, score int) error {
	err := s.checkConnection()
	if err != nil {
		return err
	}

	err = s.conn.QueryRow(
		context.Background(),
		`update assignment set score=$1 where assignmentid=$2;`,
		score, assignmentid,
	).Scan()
	if err == pgx.ErrNoRows {
		return nil
	}
	return nil
}

func (s *service) GetAllAssignment(username string) ([]Assignment, error) {
	err := s.checkConnection()
	if err != nil {
		return nil, err
	}

	rows, _ := s.conn.Query(
		context.Background(),
		`select * from assignment where createdby=$1;`,
		username,
	)

	var assignments []Assignment
	for rows.Next() {
		var a Assignment
		err := rows.Scan(&a.Id, &a.gradingfile, &a.createdBy, &a.totalScore)
		if err != nil {
			return nil, err
		}
		assignments = append(assignments, a)
	}
	return assignments, rows.Err()
}

func (s *service) RegisterSubmission(assignmentid, username string) (string, error) {
	err := s.checkConnection()
	if err != nil {
		return "", err
	}

	var fileid int
	err = s.conn.QueryRow(
		context.Background(),
		`insert into submission 
    (username, assignment) 
      values ($1, $2) 
    returning fileid;`,
		username, assignmentid).
		Scan(&fileid)
	return strconv.Itoa(fileid), err
}

func (s *service) ViewSubmissionByUser(username string) ([]Submission, error) {
	err := s.checkConnection()
	if err != nil {
		return nil, err
	}

	rows, _ := s.conn.Query(context.Background(),
		`
    select submissionid, assignment, isgraded, obtainedscore, feedback 
    from submission 
    where username=$1
    ;
    `, username)

	var results []Submission

	for rows.Next() {
		result := Submission{
			Username: username,
		}

		err := rows.Scan(
			&result.SubmissionID,
			&result.AssignmentID,
			&result.Isgraded,
			&result.Obtainedscore,
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

func (s *service) ViewSubmissionByAssignment(assignmentid string) ([]Submission, error) {
	err := s.checkConnection()
	if err != nil {
		return nil, err
	}

	rows, _ := s.conn.Query(context.Background(),
		`
    select submissionid, username, isgraded, obtainedscore, feedback 
    from submission 
    where username=$1
    ;
    `, assignmentid)

	var results []Submission

	for rows.Next() {
		result := Submission{
			AssignmentID: assignmentid,
		}

		err := rows.Scan(
			&result.SubmissionID,
			&result.Username,
			&result.Isgraded,
			&result.Obtainedscore,
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

// Basically gets the best submission of a user for the given assignment
// Groups by obtained score
// func getScoreRanges(assignmentid string) error

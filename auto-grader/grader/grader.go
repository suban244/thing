package grader

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Service interface {
	HandleFileUpload(filename string) error
}

type service struct {
	conn *pgx.Conn
}

type submission struct {
	filename string
	username string
	isGraded bool
}

// NewService is used to create a single instance of the service
func NewService(conn *pgx.Conn) Service {
	return &service{
		conn: conn,
	}
}

func (s *service) HandleFileUpload(filename string) error {
	s.conn.QueryRow(context.Background(), "").Scan()
	return nil
}

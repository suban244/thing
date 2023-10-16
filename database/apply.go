package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	path := "./init.sql"

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		log.Fatal("No sql file, add a init.sql in the directory where the program is")
	}
	sql := string(c)
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		log.Fatalln("Error in running SQL", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Pom struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Task  string    `json:"task"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	engine := html.New("./views/", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	// autoGrader := app.Group("auto-grader")

	app.Get("/auto-grader", func(c *fiber.Ctx) error {
		return c.Render("auto-grader", fiber.Map{})
	})

	app.Post("/auto-grader/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}
		return c.SaveFile(file, fmt.Sprintf("./uploaded-files/%s", file.Filename))
	})

	app.Get("/pom", func(c *fiber.Ctx) error {
		return c.Render("pom", fiber.Map{})
	})

	app.Post("/pom/api/addPom", func(c *fiber.Ctx) error {
		var data Pom

		if err := c.BodyParser(&data); err != nil {
			fmt.Print(err)
			return err
		}
		fmt.Println(data)

		return c.Status(200).JSON(&fiber.Map{
			"status": "ok",
		})
	})

	port = ":" + port
	log.Fatal(app.Listen("0.0.0.0" + port))
}

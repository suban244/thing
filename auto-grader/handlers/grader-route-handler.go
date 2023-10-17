package handlers

import (
	"fmt"
	"log"
	"thing/auto-grader/grader"

	"github.com/gofiber/fiber/v2"
)

func Index() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("auto-grader", fiber.Map{})
	}
}

func UploadFile(g grader.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			log.Println(err)
			return err
		}

		payload := struct {
			Username string `json:"username"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		// TODO: Check file extention

		// TODO: Upload Data to a database
		fileid, err := g.RegisterFile(payload.Username, file.Filename)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = c.SaveFile(file, fmt.Sprintf("./uploaded-files/%s", fileid))
		if err != nil {
			log.Println(err)
			return err
		}

		// TODO: Launch a process to grade file
		g.LaunchGrader(fileid)

		return c.Render("file-upload-success", fiber.Map{})
	}
}

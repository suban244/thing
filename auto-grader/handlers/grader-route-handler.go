package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Index() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("auto-grader", fiber.Map{})
	}
}

func UploadFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}

		// TODO: Check file extention

		// TODO: Handle dublicate names
		err = c.SaveFile(file, fmt.Sprintf("./uploaded-files/%s", file.Filename))
		if err != nil {
			return err
		}

		// A service handles this

		// TODO: Upload Data to a database

		// TODO: Launch a process to grade file

		return c.Render("file-upload-success", fiber.Map{})
	}
}

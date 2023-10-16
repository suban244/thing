package autoGrader

import (
	"thing/auto-grader/handlers"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router) {
	app.Get("/", handlers.Index())
	app.Post("/upload", handlers.UploadFile())
}

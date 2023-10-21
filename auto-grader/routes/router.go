package graderroutes

import (
	"thing/auto-grader/grader"
	"thing/auto-grader/handlers"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router, g grader.Service) {
	app.Get("/", handlers.Index())
	app.Get("/view", handlers.ViewStatus())
	app.Post("/upload", handlers.UploadFile(g))
	app.Post("/getSerachResult", handlers.ReturnResult(g))
}

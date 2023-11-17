package graderroutes

import (
	"github.com/gofiber/fiber/v2"

	"thing/auto-grader/grader"
	"thing/auto-grader/handlers"
)

func Router(app fiber.Router, g grader.Service) {
	app.Get("/", handlers.Index())
	app.Get("/view", handlers.ViewStatus())
	app.Post("/upload", handlers.UploadFile(g))
	app.Post("/getSerachResult", handlers.ReturnResult(g))

	app.Get("/assignment", handlers.ViewAssignments())
  app.Post("/createNewAssignment",handlers.CreateNewAssignment())
}

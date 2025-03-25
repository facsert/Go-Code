package router

import (
	"github.com/gofiber/fiber/v3"

	"devops/api/v1/board"
	// "fibert/api/v1/scan"
)

func Init(app *fiber.App) {
	api := app.Group("/api/v1")
    
    board.Init(api)
    // article.Init(api)
	// scan.Init(api)
}
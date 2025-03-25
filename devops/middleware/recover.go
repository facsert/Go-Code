package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func RecoverInit(app *fiber.App) {
	app.Use(recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: false,
		// StackTraceHandler: defaultStackTraceHandler,
	}))

	// Default
	// return recover.New(recover.Config{
	// 	Next:              nil,
	// 	EnableStackTrace:  false,
	// 	StackTraceHandler: defaultStackTraceHandler,
	// })
}
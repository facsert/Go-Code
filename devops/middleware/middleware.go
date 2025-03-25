package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func Init(app *fiber.App) {
	CorsInit(app)
    LoggerInit(app)
	RecoverInit(app)
}
package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func CorsInit(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     []string{"*"},
		AllowMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		},
		AllowHeaders:        []string{"*"},
		AllowCredentials:    false,
		ExposeHeaders:       []string{},
		MaxAge:              0,
		AllowPrivateNetwork: false,
	}))

	// Default
	// return cors.New(cors.Config{
	// 	Next:             nil,
	// 	AllowOriginsFunc: nil,
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods: []string{
	// 		fiber.MethodGet,
	// 		fiber.MethodPost,
	// 		fiber.MethodHead,
	// 		fiber.MethodPut,
	// 		fiber.MethodDelete,
	// 		fiber.MethodPatch,
	// 	},
	// 	AllowHeaders:        []string{},
	// 	AllowCredentials:    false,
	// 	ExposeHeaders:       []string{},
	// 	MaxAge:              0,
	// 	AllowPrivateNetwork: false,
	// })
}
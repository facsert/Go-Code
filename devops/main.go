package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"

	"devops/middleware"
	"devops/utils/comm"
	"devops/utils/database"
	"devops/utils/router"
)

var (
	host = "0.0.0.0"
	port = 3000
	app *fiber.App
)



func main() {
    
	app = NewApp()

	comm.Init()
	middleware.Init(app)
	database.Init()
	router.Init(app)

    Listen(fmt.Sprintf("%v:%v", host, port))
}

func NewApp() *fiber.App {
	return fiber.New(fiber.Config{
		ServerHeader:  "Fiber",
		AppName: "Test App v1.0.1",

		CaseSensitive: true,
		StrictRouting: false,

		BodyLimit: 4 * 1024 * 1024,
	})
}

func Listen(addr string) {
	log.Fatal(app.Listen(addr, fiber.ListenConfig{
		EnablePrefork: true,
		DisableStartupMessage: false,

		EnablePrintRoutes: true,
	}))
}
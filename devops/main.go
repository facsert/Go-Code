package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
    
	"devops/middleware"
	"devops/pkg/comm"
	"devops/pkg/router"
)

const (
	HOST = "0.0.0.0"
	PORT = 3100
)

func main() {
	app := NewApp(HOST, PORT)
	app.Init()
    app.Listen()
}

type Server struct {
	Host string
	Port int
    App *fiber.App
}

func NewApp(host string, port int) *Server {
	return &Server{
		Host: host,
		Port: port,
		App: fiber.New(fiber.Config{
			ServerHeader: "Schedule Fiber",
			AppName: "Devops v1.0.0",
			CaseSensitive: true,
			StrictRouting: false,
			BodyLimit: 4 * 1024 * 1024,
		}),
	}
}

func (s *Server) Init() {
	comm.Init()
	middleware.Init(s.App)
	router.Init(s.App)
}

func (s *Server) Listen() {
	log.Fatal(s.App.Listen(fmt.Sprintf("%s:%d", HOST, PORT), fiber.ListenConfig{
		EnablePrefork: true,
		DisableStartupMessage: false,
		EnablePrintRoutes: true,
	}))
}
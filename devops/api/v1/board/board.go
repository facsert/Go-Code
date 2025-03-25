package board

import (
	// "fmt"

	"github.com/gofiber/fiber/v3"
	// "github.com/gofiber/fiber/v3/log"
	
	// "devops/utils/database"
)

func Init(api fiber.Router) {
    router := api.Group("/boards")

	router.Get("/", GetBoards)
	router.Get("/:id", GetBoard)
	router.Post("/insert", CreateBoard)
	router.Post("/update", UpdateBoard)
	router.Post("/delete/:id", DeleteBoard)
}

func GetBoards(c fiber.Ctx) error {
    return c.SendString("get boards")
}

func GetBoard(c fiber.Ctx) error {
    return c.SendString("get board" + c.Params("id"))
}

func CreateBoard(c fiber.Ctx) error {
    return c.SendString("create board")
}

func UpdateBoard(c fiber.Ctx) error {
    return c.SendString("update board")
}

func DeleteBoard(c fiber.Ctx) error {
    return c.SendString("delete board" + c.Params("id"))
}
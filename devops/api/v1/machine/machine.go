package machine

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"

	"devops/pkg/database"
	"devops/pkg/req"
)

var (
	TABLE_MACHINE = "machines"
)


func Init(api fiber.Router) {
	r := api.Group("/machine")

	r.Get("/", GetMachines)
	r.Get("/:id", GetMachine)
	r.Post("/insert", InsertMachine)
	r.Post("/update/:id", UpdateMachine)
	r.Post("/delete/:id", DeleteMachine)
}


func GetMachines(c fiber.Ctx) error {
	cur := database.NewDB()
	var machines []Machine
	if err := cur.Where("deleted = false").Find(&machines).Error; err != nil {
		log.Errorf("select boards failed: %v\n", err)
		return c.JSON(req.NewErrJson([]Machine{}, fmt.Sprintf("select failed: %v\n", err)))
	}
	return c.JSON(req.NewPassJson(machines, ""))
}

func GetMachine(c fiber.Ctx) error {
	cur, id := database.NewDB(), c.Params("id")
	var machine Machine
	if _, err := strconv.Atoi(id); err != nil {
		log.Errorf("id %s error: %v", id, err)
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("id %s error: %v\n", id, err)))
	}
	if err := cur.Where("id = ?", id).First(&machine).Error; err != nil {
		log.Errorf("select machine failed: %v\n", err)
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("select failed: %v\n", err)))
	}
	return c.JSON(req.NewPassJson(machine, ""))
}

func InsertMachine(c fiber.Ctx) error {
	m := new(Machine)
	if err := c.Bind().JSON(&m); err != nil {
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("parse json failed: %v\n", err)))
	}
	cur := database.NewDB()
	if err := cur.Create(&m).Error; err != nil {
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("insert data failed: %v\n", err)))
	}
    return c.JSON(req.NewPassJson(nil, "Add 1"))
}

func UpdateMachine(c fiber.Ctx) error {
	m := new(Machine)
	if err := c.Bind().JSON(&m); err != nil {
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("parse json failed: %v\n", err)))
	}

	cur, id := database.NewDB(), c.Params("id")
	if _, err := strconv.Atoi(id); err != nil {
		log.Errorf("id %s error: %v", id, err)
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("id %s error: %v\n", id, err)))
	}

	if err := cur.Where("id = ?", id).Updates(m).Error; err != nil {
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("update data failed: %v\n", err)))
	}
	return c.JSON(req.NewPassJson(nil, "Update 1"))
}

func DeleteMachine(c fiber.Ctx) error {
	m := new(Machine)
	if err := c.Bind().JSON(&m); err != nil {
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("parse json failed: %v\n", err)))
	}
	id := c.Params("id")
	if _, err := strconv.Atoi(id); err != nil {
		log.Errorf("id %s error: %v", id, err)
		return c.JSON(req.NewErrJson(nil, fmt.Sprintf("id %s error: %v\n", id, err)))
	}
	cur := database.NewDB()
	cur.Where("id = ?", id).Update("delete", true)
	return c.JSON(req.NewPassJson(nil, "Delete 1"))
}
package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello, world!")
	ConnectDatabase()
	CreateTable()

	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	app.Listen(":8080")
}

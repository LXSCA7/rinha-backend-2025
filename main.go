package main

import (
	"encoding/json"
	"fmt"

	"github.com/LXSCA7/rinha-backend-2025/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello, world!")
	InitDatabase()

	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})

	app.Post("/payments", func(c *fiber.Ctx) error {
		var req models.Payment
		body := c.Body()
		err := json.Unmarshal(body, &req)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"error":      err.Error(),
				"expected":   req,
			})
		}

		InsertTable(req.CorrelationId, req.Amount, "default", "processed")
		return c.SendStatus(fiber.StatusAccepted)
	})

	app.Listen(":8080")
	// InitRabbitMq()
}

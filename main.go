package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("/", Hello)
	log.Print("app is listening on 8080")
	app.Listen(":8080")
}

func Hello(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "go deployment check for go series",
	})
}

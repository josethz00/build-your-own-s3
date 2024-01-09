package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/bucket", func(c fiber.Ctx) error {
		return c.Status(201).JSON(fiber.Map{
			"message": "BUCKET CREATED SUCCESFULLY",
		})
	})

	app.Get("/buckets", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "BUCKETS LIST",
		})
	})

	app.Listen(":8778")
}

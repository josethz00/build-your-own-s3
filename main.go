package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/josethz00/build-your-own-s3/utils"
)

func main() {
	app := fiber.New()

	app.Use(func(c fiber.Ctx) error {
		accessKey := c.Get("x-access-key")
		secretAccessKey := c.Get("x-secret-access-key")

		if accessKey == "" || secretAccessKey == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "INVALID ACCESS KEY OR SECRET ACCESS KEY",
			})
		}

		if utils.CheckApiCredentials(accessKey, secretAccessKey) {
			return c.Next()
		}

		return c.Status(401).JSON(fiber.Map{
			"message": "INVALID ACCESS KEY OR SECRET ACCESS KEY",
		})
	})

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

	app.Get("/bucket/:id", func(c fiber.Ctx) error {
		bucketID := c.Params("id")
		return c.Status(200).JSON(fiber.Map{
			"message": "BUCKET DETAILS",
			"id":      bucketID,
		})
	})

	app.Listen(":8778")
}

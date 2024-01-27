package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/josethz00/build-your-own-s3/utils"
)

type CreateBucketRequest struct {
	Name   string `json:"name"` // this json:"name" is called a tag and it's used to map the json key to the struct field
	Public bool   `json:"public"`
}

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/bucket", func(c *fiber.Ctx) error {
		bucket := new(CreateBucketRequest)

		if err := c.BodyParser(bucket); err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(500)
		}

		fmt.Println("bucket = ", bucket)

		return c.Status(201).JSON(fiber.Map{
			"message": "BUCKET CREATED SUCCESFULLY",
		})
	})

	app.Get("/buckets", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "BUCKETS LIST",
		})
	})

	app.Get("/bucket/:id", func(c *fiber.Ctx) error {
		bucketID := c.Params("id")
		return c.Status(200).JSON(fiber.Map{
			"message": "BUCKET DETAILS",
			"id":      bucketID,
		})
	})

	app.Listen(":8778")
}

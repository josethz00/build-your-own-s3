package main

import (
	"context"
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/josethz00/build-your-own-s3/internal/db"
	"github.com/josethz00/build-your-own-s3/utils"
)

type CreateBucketRequest struct {
	Name   string `json:"name"` // this json:"name" is called a tag and it's used to map the json key to the struct field
	Public bool   `json:"public"`
}

func main() {
	app := fiber.New()

	// This is a Go Context, it's used to run concurrent/background operations
	ctx := context.Background()

	// Use pgx
	dbconn, err := pgx.Connect(ctx, "user=s3db host=localhost port=26257 dbname=s3db sslmode=disable")
	if err != nil {
		fmt.Println("Failed to parse PGX config:", err)
		return
	}

	// Close the connection before the main function ends
	defer dbconn.Close(ctx)

	queries := db.New(dbconn)

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	app.Use(func(c *fiber.Ctx) error {
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
			"message":  "BUCKET CREATED SUCCESFULLY",
			"bucketID": node.Generate().String(),
		})
	})

	app.Get("/buckets", func(c *fiber.Ctx) error {
		buckets, err := queries.ListBuckets(ctx)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "FAILED TO LIST BUCKETS",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"buckets": buckets,
		})
	})

	app.Get("/bucket/:id", func(c *fiber.Ctx) error {
		bucketID := c.Params("id")
		return c.Status(200).JSON(fiber.Map{
			"message": "BUCKET DETAILS",
			"id":      bucketID,
		})
	})

	startappErr := app.Listen(":8778")

	if startappErr != nil {
		panic(startappErr)
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var App *fiber.App

var Ctx *fiber.Ctx

var rdb *redis.Client

var ctx_redis context.Context

func main() {

	fmt.Println("11111")

	//	var	ctx_redis context.Context

	ctx_redis = context.Background()

	// 1. Connect to Redis server
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "sucesso123",
		DB:       0,
	})

	fmt.Println("22222")

	App = fiber.New()

	App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	App.Get("/", HelloWorld)

	// 2. Set a key with NO expiration (0)
	/*
		err := rdb.Set(ctx, "user_name", "Fabian", 0).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully set 'user_name'")
	*/

	// 3. Set a key WITH an expiration (TTL) of 10 seconds
	/*
		err = rdb.Set(ctx, "session_token", "xyz123", 10*time.Second).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully set 'session_token' with a 10s TTL")
	*/

	// 4. Retrieve the value using Get

	App.Listen(":8071")

}

func HelloWorld(c *fiber.Ctx) error {

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	//	return c.Status(fiber.StatusOK).SendString(" OK *** OK *** Ok  **  OK ***")

	val, err := rdb.Get(ctx_redis, "user_name").Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusOK).SendString("Key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		return c.Status(fiber.StatusOK).SendString("Retrieved value: " + val)
	}

}

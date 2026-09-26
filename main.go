package main

import (
	"context"
	"fmt"
	"log"

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
		Addr:     "leander.vps-kinghost.net:6379",
		Password: "sucesso123",
		DB:       0,
	})

	//		Addr:     "127.0.0.1:6379",
	//	Password: "sucesso123",

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

	// Store JSON document using JSONSet
	doc := `{
        "name": "Alice",
        "age": 30,
        "email": "alice@example.com",
        "roles": ["admin", "user"],
        "address": {"city": "Portland", "zip": "97201"}
    }`

	err := rdb.Do(ctx_redis, "JSON.SET", "aquivaiip2", "$", doc).Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stored")

	// Get full document
	result, err := rdb.Do(ctx_redis, "JSON.GET", "aquivaiip2").Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Get nested field using JSONPath
	name, err := rdb.Do(ctx_redis, "JSON.GET", "aquivaiip2", "$.name").Text()
	fmt.Println(name) // ["Alice"]

	city, err := rdb.Do(ctx_redis, "JSON.GET", "aquivaiip2", "$.address.city").Text()
	fmt.Println(city) // ["Portland"]

	/*



		val, err := rdb.Get(ctx_redis, "user_name").Result()
		if err == redis.Nil {

			err := rdb.Set(ctx_redis, "user_name", "Fabian", 0).Err()
			if err != nil {
				panic(err)
			}
			return c.Status(fiber.StatusOK).SendString("Successfully set 'user_name'")

		} else if err != nil {
			panic(err)
		} else {
			return c.Status(fiber.StatusOK).SendString("Retrieved value: " + val + "--" + greeting)
		}

	*/

	return c.Status(fiber.StatusOK).SendString("Nome: " + name + "  Cidade: " + city)

}

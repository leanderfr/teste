package main

import (
	"fmt"
	"log"

	//	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var App *fiber.App

var Ctx *fiber.Ctx

/*
var rdb *redis.Client

var ctx_redis context.Context
*/
func main() {

	fmt.Println("11111__")
	App = fiber.New()

	App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	App.Get("/", HelloWorld)

	//App.Listen(":8071")

	//fmt.Println("22222")

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("rodando https")
		fmt.Println("rodando https 222")

		return c.SendString("Secure connection")
	})
	log.Fatal(app.ListenTLS(":8443", "./fullchain.pem", "./privkey.key"))

}

func HelloWorld(c *fiber.Ctx) error {

	fmt.Println("2222")

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.Status(fiber.StatusOK).SendString(" OK *** OK *** Ok  **  OK ***")

	return nil

}

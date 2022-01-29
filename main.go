package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


//context is whatever that comes with your request
func welcome(c *fiber.Ctx)error  {
	return c.SendString("Welcome to my awesome API")
}

func main() {

	//creating an instance of my fiber app
	app := fiber.New()
	
	//Route
	app.Get("/api",welcome)

	//prints up the current state of our server
	log.Fatal(app.Listen(":3000"))


}
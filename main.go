package main

import (
	"log"

	"github.com/MbuguaCaleb/go_rest_api_one/database"
	"github.com/MbuguaCaleb/go_rest_api_one/routes"
	"github.com/gofiber/fiber/v2"
)

//context is whatever that comes with your request
func welcome(c *fiber.Ctx)error  {
	return c.SendString("Welcome to my awesome API")
}

//function that returns all of our Routes
func setUpRoutes(app *fiber.App){
	//Welcome endpoint
	app.Get("/api",welcome)

	//User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id",routes.UpdateUser)
	app.Delete("/api/users/:id",routes.DeleteUser)

	//Product Endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id",routes.UpdateProduct)
	app.Delete("/api/products/:id",routes.DeleteProduct)

	


}

func main() {

	database.ConnectDb()
	//creating an instance of my fiber app
	app := fiber.New()
	
	//Calling the router function
	setUpRoutes(app)

	//prints up the current state of our server
	log.Fatal(app.Listen(":3000"))


}
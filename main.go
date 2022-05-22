package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/standingdreams/fiber-api/controllers"
)

func sneakerHandler(api fiber.Router) {
	api.Get("/", controllers.SneakerController)
}

func userHandler(api fiber.Router) {
	api.Get("/", controllers.UserController)
}

func storeHandler(api fiber.Router) {
	api.Get("/", controllers.StoreController)
}

func setUpRoutes(app *fiber.App) {

	// Main API route
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Individual API Routes
	v1.Route("/sneaker", sneakerHandler)
	v1.Route("/user", userHandler)
	v1.Route("/store", storeHandler)
}

func main() {
	app := fiber.New()

	setUpRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/standingdreams/fiber-api/sneakerController"
	"github.com/standingdreams/fiber-api/storeController"
	"github.com/standingdreams/fiber-api/userController"
)

func sneakerHandler(api fiber.Router) {
	api.Get("/", sneakerController.Controller)
}

func userHandler(api fiber.Router) {
	api.Get("/", userController.Controller)
}

func storeHandler(api fiber.Router) {
	api.Get("/", storeController.Controller)
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

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/standingdreams/fiber-api/shared"
)

func StoreController(c *fiber.Ctx) error {
	return c.SendString(shared.GenerateString("store controller"))
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/standingdreams/fiber-api/shared"
)

func SneakerController(c *fiber.Ctx) error {
	return c.SendString(shared.GenerateString("sneakerController"))
}

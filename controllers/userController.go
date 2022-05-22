package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/standingdreams/fiber-api/shared"
)

func UserController(c *fiber.Ctx) error {
	return c.SendString(shared.GenerateString("user controller"))
}

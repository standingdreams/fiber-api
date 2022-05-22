package sneakerController

import "github.com/gofiber/fiber/v2"

func generateString(title string) string {
	returnValue := "API Route: " + title
	return returnValue
}

func Controller(c *fiber.Ctx) error {
	return c.SendString(generateString("sneakerController"))
}

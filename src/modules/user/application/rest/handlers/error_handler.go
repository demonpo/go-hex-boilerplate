package handlers

import "github.com/gofiber/fiber/v2"

func HandleError(ctx *fiber.Ctx, statusCode int, err error) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"error": err.Error(),
	})
}

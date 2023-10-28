package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AleksMedovnik/Teach-student__personal-account/models"
)

func GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

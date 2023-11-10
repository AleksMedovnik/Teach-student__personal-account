package controllers

import (
	"github.com/AleksMedovnik/Teach-student__personal-account/initializers"
	"github.com/AleksMedovnik/Teach-student__personal-account/models"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	initializers.DB.Find(&users)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"users": users}})
}
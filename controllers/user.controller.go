package controllers

import (
	"errors"

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

func findUser(id int, user *models.User) error {
	initializers.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}

func GetOwnGroups(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)
	var userSpec models.User
	groups := &user.Groups
	if err := findUser(int(user.ID), &userSpec); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	err := initializers.DB.Model(&userSpec).Association("Groups").Find(&groups)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"groups": groups}})
}
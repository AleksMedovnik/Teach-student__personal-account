package controllers

import (
	"errors"
	"strings"

	"github.com/AleksMedovnik/Teach-student__personal-account/initializers"
	"github.com/AleksMedovnik/Teach-student__personal-account/models"
	"github.com/gofiber/fiber/v2"
)

func CreateGroup(c *fiber.Ctx) error {
	var payload *models.GroupInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

	}

	newGroup := models.Group{
		Number: payload.Number,
		Users:  payload.Users,
	}

	result := initializers.DB.Create(&newGroup)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Group with that number already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"group": models.FilterGroupResord(&newGroup)}})
}

func GetGroups(c *fiber.Ctx) error {
	groups := []models.Group{}
	initializers.DB.Find(&groups)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"groups": groups}})
}

func findGroup(id int, group *models.Group) error {
	initializers.DB.Find(&group, "id = ?", id)
	if group.ID == 0 {
		return errors.New("group does not exist")
	}
	return nil
}

func GetGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var group models.Group
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findGroup(id, &group); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"group": group}})
}

func UpdateGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var group models.Group
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findGroup(id, &group); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateGroup struct {
		Number uint16 `json:"number" validate:"required"`
	}
	var updateData UpdateGroup
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	group.Number = updateData.Number
	initializers.DB.Save(&group)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"group": group}})
}

func DeleteGroup(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var group models.Group
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findGroup(id, &group); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Delete(&group)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"group": group}})
}

func GetGroupUsers(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var group models.Group
	users := &group.Users
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findGroup(id, &group); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	err = initializers.DB.Model(&group).Association("Users").Find(&users)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"users": users}})
}

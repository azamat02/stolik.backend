package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mborders/logmatic"
	"stolik.online/database"
	"stolik.online/models"
)

var logger = logmatic.NewLogger()

func CheckPhone(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	user := models.User{}

	result := database.DB.Where("phone=?", data["phone"]).First(&user)

	if result.Error != nil && result.Error.Error() != "record not found" {
		c.Status(fiber.StatusInternalServerError)
		logger.Error("ERROR: %s", result.Error)
		return c.JSON(fiber.Map{
			"message": "unable to check user",
		})
	}

	if result.RowsAffected == 0 {
		c.Status(fiber.StatusNotFound)
		logger.Warn("user with phone: %s not found", data["phone"])
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(user)
}

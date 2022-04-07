package controllers

import (
	"TTU_CORE_ERP_PORTAL/database"
	"TTU_CORE_ERP_PORTAL/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}

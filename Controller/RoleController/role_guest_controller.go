package RoleController

import (
	database "portalapp/Database"
	userRole "portalapp/Model/UserModel"

	"github.com/gofiber/fiber/v2"
)

func GetAllRole(ctx *fiber.Ctx) error {
	var roles []userRole.Role
	database.DB.Find(&roles)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"roles": roles})
}

func CreateRole(ctx *fiber.Ctx) error {

	var body struct {
		Name        string `json:"name"`
		Description string `json:"desription"`
	}

	err := ctx.BodyParser(&body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "InvalidData"})
	}

	roleObj := userRole.Role{Name: body.Name, Description: body.Description}
	result := database.DB.Create(&roleObj)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create user role."})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "message": "Sucessfully created a user role."})
}

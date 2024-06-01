package UserController

import (
	"portalapp/Database"
	usermodel "portalapp/Model/UserModel"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(ctx *fiber.Ctx) error {
	var users []usermodel.User
	Database.DB.Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"users": users})
}



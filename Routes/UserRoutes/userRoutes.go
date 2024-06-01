package UserRoutes

import (
	roleController "portalapp/Controller/RoleController"
	userController "portalapp/Controller/UserController"

	"github.com/gofiber/fiber/v2"
)

func UserAuthRoute(app *fiber.App) {
	user_route := app.Group("/user")
	user_obj := user_route.Group("/")
	user_obj.Post("/addUser", userController.AddUser)
	user_obj.Post("/login", userController.Login)
	user_obj.Post("/roles", roleController.CreateRole)
}

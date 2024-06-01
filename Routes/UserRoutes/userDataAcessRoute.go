package UserRoutes

import (
	userController "portalapp/Controller/UserController"
	middleware "portalapp/Middleware/AuthMiddleware"

	"github.com/gofiber/fiber/v2"
)

func UserAccessRoute(app *fiber.App) {
	user_route := app.Group("/users")
	user_obj := user_route.Group("/", middleware.RequireAuth)
	user_obj.Get("/allUser", userController.GetAllUser)
}

package UserRoutes

import (
	userController "portalapp/Controller/UserController"

	"github.com/gofiber/fiber/v2"
)

func UserWebRoutes(app *fiber.App) {
	user_route := app.Group("/web")
	user_obj := user_route.Group("/")
	user_obj.Post("/addUser", userController.ProcessLogInData)
	user_obj.Get("/login", userController.GetLogInForm)
}

package WebUI

import (
	jobController "portalapp/Controller/JenkinsJobController"
	userController "portalapp/Controller/UserController"

	"github.com/gofiber/fiber/v2"
)

func WebUIRoute(app *fiber.App) {
	web_route := app.Group("/webui")
	web := web_route.Group("/")
	web.Get("/login", userController.GetLogInForm)
	web.Get("/allJobs", jobController.GetAllJenkinsJob)
}

package JenkinsJobRoutes

import (
	jenkinsController "portalapp/Controller/JenkinsJobController"
	middleware "portalapp/Middleware/AuthMiddleware"

	"github.com/gofiber/fiber/v2"
)

func JenkinJobRoute(app *fiber.App) {
	jenkins_route := app.Group("/jenkins")
	jenkins_job := jenkins_route.Group("/", middleware.RequireAuth)
	jenkins_job.Post("/addJob", jenkinsController.CreateJenkinsJob)
	jenkins_job.Get("/allJobs", jenkinsController.GetAllJobs)
}

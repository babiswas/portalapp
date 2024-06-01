package TestRoutes

import (
	testController "portalapp/Controller/TestController"
	middleware "portalapp/Middleware/AuthMiddleware"

	"github.com/gofiber/fiber/v2"
)

func TestRoute(app *fiber.App) {
	test_route := app.Group("/testcase")
	test_obj := test_route.Group("/", middleware.RequireAuth)
	test_obj.Post("/addTest", testController.CreateTestCase)
	test_obj.Get("/allTest", testController.GetAllTestCases)
}

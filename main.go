package main

import (
	"fmt"
	"log"
	jnknroutes "portalapp/Routes/JenkinsJobRoutes"
	testroutes "portalapp/Routes/TestRoutes"
	userroutes "portalapp/Routes/UserRoutes"
	webuiroutes "portalapp/Routes/WebUI"

	dbutil "portalapp/Database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	fmt.Println("Executing the init function block:")
	dbutil.LoadENVVar()
	dbutil.ConnectDB()
	dbutil.SyncDatabase()
}

func main() {
	log.Println("Starting the application:")
	engine := html.New("./Views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Static("webui/Static", "./Static")

	jnknroutes.JenkinJobRoute(app)
	testroutes.TestRoute(app)
	userroutes.UserAccessRoute(app)
	userroutes.UserAuthRoute(app)
	userroutes.UserWebRoutes(app)

	webuiroutes.WebUIRoute(app)
	app.Listen(":3000")
}

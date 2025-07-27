package app

import (
	"log"

	db "github.com/farhanaltariq/fiberplate/app/database"
	"github.com/farhanaltariq/fiberplate/app/middleware"
	"github.com/farhanaltariq/fiberplate/app/routes"
	utils "github.com/farhanaltariq/fiberplate/app/utils"
	_ "github.com/farhanaltariq/fiberplate/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
)

var BaseUrl = utils.GetEnv("BASE_URL", "localhost:3000")

func SetupApp() *fiber.App {
	app := fiber.New(
		fiber.Config{
			Prefork:           false,
			ReduceMemoryUsage: true,
		},
	)

	app.Use(recover.New())
	app.Get("api/swagger/*", swagger.HandlerDefault)
	app.Use(middleware.CommonMiddleware)

	routes.Init(app)

	logrus.Infoln("Server running on ", BaseUrl)

	return app
}

// @title Fiber Boilerplate API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /api
// @schemes http https
// @SecurityDefinitions.apiKey  Authorization
// @in header
// @name Authorization
func RunServer() {
	utils.CustomFormatter()
	utils.RenameBaseUrlSwagger(BaseUrl)

	if err := db.Connect(); err != nil {
		log.Fatal("Error connecting to database", err)
	}
	logrus.Infoln("Connected to database")

	app := SetupApp()

	go func() {
		if err := app.Listen(BaseUrl); err != nil {
			logrus.Errorln(err)
		}
	}()

	select {}
}

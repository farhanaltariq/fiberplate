package app

import (
	"log"

	db "github.com/farhanaltariq/fiberplate/app/database"
	"github.com/farhanaltariq/fiberplate/app/middleware"
	"github.com/farhanaltariq/fiberplate/app/routes"
	utils "github.com/farhanaltariq/fiberplate/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app.Use(middleware.CommonMiddleware)
	app.Use(middleware.AuthInterceptor)

	routes.Init(app)

	logrus.Infoln("Server running on ", BaseUrl)

	return app
}

func RunServer() {
	utils.CustomFormatter()

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

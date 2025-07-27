package routes

import (
	"github.com/farhanaltariq/fiberplate/app/controllers"
	"github.com/farhanaltariq/fiberplate/app/database"
	"github.com/farhanaltariq/fiberplate/app/middleware"
	"github.com/farhanaltariq/fiberplate/app/services"
	"github.com/gofiber/fiber/v2"
)

func initServices() middleware.Services {
	db := database.GetDBConnection()
	return middleware.Services{
		DB:          db,
		AuthService: services.NewAuthService(db),
		UserService: services.NewUserService(db),
	}
}

func Init(app *fiber.App) {
	services := initServices()

	api := app.Group("/api")
	api.Get("/", controllers.NewMiscController(services).HealthCheck)

	Authentications(api.Group("/auth"), services)

	api.Use(middleware.AuthInterceptor)
	User(api.Group("/user"), services)
}

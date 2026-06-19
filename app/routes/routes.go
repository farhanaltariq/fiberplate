package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/farhanaltariq/fiberplate/app/common"
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
	// Override Huma's default error generator to use our standard ResponseMessage.
	huma.NewError = func(status int, message string, errs ...error) huma.StatusError {
		return &common.ResponseMessage{
			IsError: true,
			Code:    status,
			Message: message,
		}
	}

	services := initServices()

	// Configure Huma API
	config := huma.DefaultConfig("Fiber Boilerplate API", "1.0.0")
	config.DocsPath = "/api/swagger"
	config.OpenAPIPath = "/api/openapi"

	// Add Bearer Auth security scheme
	config.Components = &huma.Components{
		SecuritySchemes: map[string]*huma.SecurityScheme{
			"bearerAuth": {
				Type:         "http",
				Scheme:       "bearer",
				BearerFormat: "JWT",
				Description:  "Enter your JWT access token",
			},
		},
	}

	api := humafiber.New(app, config)

	// Register Health Check
	miscController := controllers.NewMiscController(services)
	huma.Register(api, huma.Operation{
		Method:      "GET",
		Path:        "/api",
		Summary:     "Health Check",
		Description: "Check the status of the server",
		Tags:        []string{"Misc"},
	}, miscController.HealthCheck)

	Authentications(api, services)
	User(api, services)
}

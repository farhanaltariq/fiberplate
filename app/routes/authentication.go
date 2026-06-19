package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/farhanaltariq/fiberplate/app/controllers"
	"github.com/farhanaltariq/fiberplate/app/middleware"
)

func Authentications(api huma.API, service middleware.Services) {
	authController := controllers.NewAuthController(service)

	huma.Register(api, huma.Operation{
		Method:      "POST",
		Path:        "/api/auth/register",
		Summary:     "Register a new user",
		Description: "Create a new user account with username, email, country, and password",
		Tags:        []string{"Authentication"},
	}, authController.Register)

	huma.Register(api, huma.Operation{
		Method:      "POST",
		Path:        "/api/auth/login",
		Summary:     "Login",
		Description: "Authenticate using username or email and password to receive an access token",
		Tags:        []string{"Authentication"},
	}, authController.Login)
}

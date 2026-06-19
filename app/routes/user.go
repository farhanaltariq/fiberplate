package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/farhanaltariq/fiberplate/app/controllers"
	"github.com/farhanaltariq/fiberplate/app/middleware"
)

func User(api huma.API, service middleware.Services) {
	userController := controllers.NewUserController(service)
	huma.Register(api, huma.Operation{
		Method:      "POST",
		Path:        "/api/user",
		Summary:     "Get List User",
		Description: "Retrieve a list of users (protected route)",
		Tags:        []string{"User"},
		Security: []map[string][]string{
			{"bearerAuth": {}},
		},
	}, userController.GetListUser)
}

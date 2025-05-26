package routes

import (
	"github.com/farhanaltariq/fiberplate/app/controllers"
	"github.com/farhanaltariq/fiberplate/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func User(router fiber.Router, service middleware.Services) {
	userController := controllers.NewUserController(service)
	router.Post("/", userController.GetListUser)
	// router.Post("/register", controllers.Register)
}

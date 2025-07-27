package controllers

import (
	"github.com/farhanaltariq/fiberplate/app/common/codes"
	"github.com/farhanaltariq/fiberplate/app/common/status"
	"github.com/farhanaltariq/fiberplate/app/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetListUser(c *fiber.Ctx) error
}

func NewUserController(service middleware.Services) UserController {
	return &controller{service}
}

func (s *controller) GetListUser(c *fiber.Ctx) error {
	return status.Success(c, codes.OK, "OK")
}

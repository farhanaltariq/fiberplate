package controllers

import (
	"context"
	"net/http"

	"github.com/farhanaltariq/fiberplate/app/common"
	"github.com/farhanaltariq/fiberplate/app/middleware"
)

type GetListUserInput struct{}

type GetListUserOutput struct {
	Body common.ResponseMessage
}

type UserController interface {
	GetListUser(ctx context.Context, input *GetListUserInput) (*GetListUserOutput, error)
}

func NewUserController(service middleware.Services) UserController {
	return &controller{service}
}

func (s *controller) GetListUser(ctx context.Context, input *GetListUserInput) (*GetListUserOutput, error) {
	return &GetListUserOutput{
		Body: common.ResponseMessage{
			IsError: false,
			Code:    http.StatusOK,
			Message: "OK",
		},
	}, nil
}

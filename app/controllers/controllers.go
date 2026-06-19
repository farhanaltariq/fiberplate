package controllers

import (
	"context"
	"net/http"

	"github.com/farhanaltariq/fiberplate/app/common"
	"github.com/farhanaltariq/fiberplate/app/middleware"
)

type HealthCheckInput struct{}

type HealthCheckOutput struct {
	Body common.ResponseMessage
}

type MiscController interface {
	HealthCheck(ctx context.Context, input *HealthCheckInput) (*HealthCheckOutput, error)
}

type controller struct {
	middleware.Services
}

func NewMiscController(service middleware.Services) MiscController {
	return &controller{service}
}

func (server *controller) HealthCheck(ctx context.Context, input *HealthCheckInput) (*HealthCheckOutput, error) {
	return &HealthCheckOutput{
		Body: common.ResponseMessage{
			IsError: false,
			Code:    http.StatusOK,
			Message: "Server Running",
		},
	}, nil
}

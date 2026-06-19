package controllers

import (
	"context"
	"net/http"

	"github.com/farhanaltariq/fiberplate/app/common"
	"github.com/farhanaltariq/fiberplate/app/common/usertype"
	"github.com/farhanaltariq/fiberplate/app/database/models"
	"github.com/farhanaltariq/fiberplate/app/middleware"
	"github.com/farhanaltariq/fiberplate/app/utils"
	"github.com/sirupsen/logrus"
)

type RegisterInput struct {
	Body models.Register
}

type RegisterOutput struct {
	Body common.ResponseMessage
}

type LoginInput struct {
	Body models.Login
}

type LoginOutput struct {
	Body models.AuthenticationResponse
}

type AuthenticationController interface {
	Register(ctx context.Context, input *RegisterInput) (*RegisterOutput, error)
	Login(ctx context.Context, input *LoginInput) (*LoginOutput, error)
}

func NewAuthController(service middleware.Services) AuthenticationController {
	return &controller{service}
}

func (s *controller) Register(ctx context.Context, input *RegisterInput) (*RegisterOutput, error) {
	auth := input.Body

	if auth.Password != auth.ConfirmPassword {
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: "Password and confirm password does not match",
		}
	}

	userOrm := models.User{
		Username: auth.Username,
		Email:    auth.Email,
		Address:  auth.Country,
		UserType: usertype.ADMIN,
	}

	data, err := s.UserService.GetDataByUsernameOrEmail(userOrm)
	if err != nil || data.ID != 0 {
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: "Username or email already used",
		}
	}

	userData, err := s.UserService.InsertOrUpdate(userOrm)
	if err != nil {
		logrus.Errorln("Failed to register user", err)
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	pass, salt := utils.Encrypt(auth.Password)
	authOrm := &models.Authentications{
		Password: pass,
		Salt:     salt,
		UserId:   userData.ID,
	}

	if err := s.AuthService.InsertOrUpdate(*authOrm); err != nil {
		logrus.Errorln("Failed to register auth", err)
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return &RegisterOutput{
		Body: common.ResponseMessage{
			IsError: false,
			Code:    http.StatusOK,
			Message: "Success",
		},
	}, nil
}

func (s *controller) Login(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	cred := input.Body

	if cred.UsernameOrEmail == "" {
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: "Username or email is required",
		}
	}

	data, err := s.UserService.GetDataByUsernameOrEmail(models.User{Username: cred.UsernameOrEmail, Email: cred.UsernameOrEmail})
	if err != nil || data == (models.User{}) {
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: "Username or email not found",
		}
	}

	authData, err := s.AuthService.GetDataByUserId(data.ID)
	if err != nil {
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	pass, err := utils.Decrypt(authData.Password, authData.Salt)
	if err != nil {
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if pass != cred.Password {
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: "Invalid credentials",
		}
	}

	token, err := utils.GenerateToken(&data)
	if err != nil {
		return nil, &common.ResponseMessage{
			IsError: true,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return &LoginOutput{
		Body: models.AuthenticationResponse{
			Status:      "Success",
			AccessToken: token,
			ExpiredAt:   utils.GetExpirationTime().Format("2006-01-02 15:04:05"),
		},
	}, nil
}

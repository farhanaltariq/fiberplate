package middleware

import (
	"strings"

	"github.com/farhanaltariq/fiberplate/app/common/status"
	"github.com/farhanaltariq/fiberplate/app/database/models"
	"github.com/farhanaltariq/fiberplate/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func CommonMiddleware(c *fiber.Ctx) error {
	endpoint := c.BaseURL() + c.Path()
	logrus.Infoln(utils.FormatMethod(c), endpoint)
	return c.Next()
}

func validateToken(c *fiber.Ctx, tokenString string, jwtSecret []byte) error {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return err
	}

	_, ok := token.Claims.(*models.Claims)
	if !ok {
		status.Error(c, fiber.ErrUnauthorized.Code, "Unauthorized")
	}

	return nil
}

func AuthInterceptor(c *fiber.Ctx) error {
	path := c.Path()
	// Skip auth for public routes: health check, authentication endpoints, and API documentation
	if path == "/api" || path == "/api/" || strings.HasPrefix(path, "/api/auth/") || strings.HasPrefix(path, "/api/swagger") || strings.HasPrefix(path, "/api/openapi") || strings.HasPrefix(path, "/docs") || strings.HasPrefix(path, "/schemas") {
		return c.Next()
	}

	jwtSecret := []byte(utils.GetEnv("JWT_SECRET", "secret"))
	authHeader := c.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return status.Error(c, fiber.ErrUnauthorized.Code, "Unauthorized")
	}

	tokenString := authHeader[7:] // Remove "Bearer " prefix

	if err := validateToken(c, tokenString, jwtSecret); err != nil {
		return err
	}

	return c.Next()
}

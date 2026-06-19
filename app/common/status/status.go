package status

import (
	"fmt"

	"github.com/farhanaltariq/fiberplate/app/common"
	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, codes int, message string, args ...any) error {
	if codes >= 200 && codes < 300 {
		return Success(c, codes, message, args...)
	}

	jsonMsg := &common.ResponseMessage{
		IsError: true,
		Code:    codes,
		Message: fmt.Sprint(message, args),
	}
	return c.Status(codes).JSON(jsonMsg)
}

func Success(c *fiber.Ctx, codes int, message string, args ...any) error {
	if !(codes >= 200) && (codes < 300) {
		return Error(c, codes, message, args...)
	}
	jsonMsg := &common.ResponseMessage{
		IsError: false,
		Code:    codes,
		Message: fmt.Sprint(message, args),
	}

	return c.Status(codes).JSON(jsonMsg)
}

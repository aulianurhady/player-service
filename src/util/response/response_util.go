package response

import (
	"github.com/gofiber/fiber/v2"
)

type restSuccess struct {
	ErrMessage string      `json:"message"`
	ErrStatus  int         `json:"status"`
	Data       interface{} `json:"data"`
}

type restErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

type restMultipleErr struct {
	ErrMessage string   `json:"message"`
	ErrStatus  int      `json:"status"`
	ErrError   []string `json:"error"`
}

func RestSuccess(c *fiber.Ctx, message string, status int, data interface{}) error {
	return c.Status(status).JSON(restSuccess{
		ErrMessage: message,
		ErrStatus:  status,
		Data:       data,
	})
}

func RestError(c *fiber.Ctx, message string, status int, err error) error {
	return c.Status(status).JSON(restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err.Error(),
	})
}

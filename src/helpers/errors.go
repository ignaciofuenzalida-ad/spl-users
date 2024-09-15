package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func InternalError(c fiber.Ctx, err error) error {
	fmt.Println(err)
	return c.
		Status(fiber.StatusInternalServerError).
		JSON(fiber.Map{"error": "Internal Server Error"})
}

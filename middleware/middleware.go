package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var IsAuthenticatedAccessToken = jwtware.New(jwtware.Config{
	SigningKey:   []byte("accessToken"),
	ErrorHandler: jwtError,
})

var IsAuthenticatedRefreshToken = jwtware.New(jwtware.Config{
	SigningKey:   []byte("refreshToken"),
	ErrorHandler: jwtError,
})

func jwtError(c *fiber.Ctx, err error) error {
	fmt.Println("hallo middleware")

	switch err.Error() {
	case "Missing or malformed JWT":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message":       "Failed",
			"Error_key":     "error_no_auth_token",
			"Error_message": "Request tidak mendapati token access pada Header",
		})
	case "token contains an invalid number of segments":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message":       "Failed",
			"Error_key":     "error_invalid_token",
			"Error_message": "Token access pada Header tidak sesuai ketentuan / settingan token",
		})
	case "Token is expired":
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message":       "Failed",
			"Error_key":     "error_expired_token",
			"Error_message": "Token access pada Header sudah kadaluarsa",
		})
	default:
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Message":       "Failed",
			"Error_key":     "error_internal_server",
			"Error_message": "Error Selain yang tercantum di sini",
		})

	}

}

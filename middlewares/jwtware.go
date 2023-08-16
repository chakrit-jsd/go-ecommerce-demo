package middlewares

import (
	"go-ecommerce-demo/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func JwtWare(userSrv ports.UserService) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("jwt.secret"))},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			str, err := token.Claims.GetIssuer()
			if err != nil {
				return fiber.ErrBadRequest
			}

			userId, err := strconv.Atoi(str)
			if err != nil {
				return fiber.ErrInternalServerError
			}

			userInfo, err := userSrv.GetUserById(&userId)
			if err != nil {
				return fiber.ErrUnauthorized
			}
			c.Locals("userInfo", userInfo)
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	})
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"matcher/pkg"
	"matcher/tools/jwt"
	"matcher/tools/zap"
	"strings"
)

func IsAuthenticated(c *fiber.Ctx) error {

	authorization := c.Get("Authorization")
	if !strings.Contains(authorization, "Bearer ") || authorization == "" || authorization == "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.NewResponse(false, pkg.ErrUnauthorized.Error(), nil))
	}

	splittedAuthorization := strings.Split(authorization, " ")
	if len(splittedAuthorization) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.NewResponse(false, pkg.ErrUnauthorized.Error(), nil))
	}

	token := splittedAuthorization[1]

	if !jwt.NewJWT(token).GetIsAuthenticated() {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.NewResponse(false, pkg.ErrUnauthorized.Error(), nil))
	}

	return c.Next()
}

func Logger(c *fiber.Ctx) error {
	zap.Logger.Info(c.String())
	return c.Next()
}

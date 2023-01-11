package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"matcher/pkg"
	"matcher/tools/jwt"
	zap_tools "matcher/tools/zap"
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

	request := map[string]string{
		"method": c.Method(),
		"url":    c.OriginalURL(),
		"body":   string(c.Body()),
		"ip":     c.IP(),
		"host":   c.Hostname(),
	}

	zap_tools.Logger.Info("Request", zap.Any("Request Object", request))

	return c.Next()
}

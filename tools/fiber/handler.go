package fiber

import (
	"github.com/gofiber/fiber/v2"
	"matcher/internal"
	"matcher/pkg"
	"strconv"
)

const driverServicePort = 8080

type IHandler interface {
	GetNearestDriver(c *fiber.Ctx) error
}

type Handler struct {
	Repository internal.IRepository
}

func NewHandler(repository internal.IRepository) IHandler {
	return &Handler{
		Repository: repository,
	}
}

// GetNearestDriver is a handler to get nearest driver
func (h *Handler) GetNearestDriver(c *fiber.Ctx) error {

	lat := c.Query("lat")
	long := c.Query("long")

	if lat == "" || long == "" {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.NewResponse(false, pkg.ErrInvalidRequest.Error(), nil))
	}

	floatLat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.NewResponse(false, pkg.ErrInvalidRequest.Error(), nil))
	}

	floatLong, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.NewResponse(false, pkg.ErrInvalidRequest.Error(), nil))
	}

	if floatLat < -90 || floatLat > 90 || floatLong < -180 || floatLong > 180 {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.NewResponse(false, pkg.ErrInvalidRequest.Error(), nil))
	}

	location, err := h.Repository.GetNearestDriver(driverServicePort, floatLat, floatLong)
	if err != nil {
		if err == pkg.ErrDriverNotFound {
			return c.Status(fiber.StatusNotFound).JSON(pkg.NewResponse(false, pkg.ErrDriverNotFound.Error(), nil))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(pkg.NewResponse(false, pkg.ErrInternalServer.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(pkg.NewResponse(true, "Driver found", location))
}

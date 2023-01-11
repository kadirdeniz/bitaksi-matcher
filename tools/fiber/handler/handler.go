package handler

import (
	"github.com/gofiber/fiber/v2"
	"matcher/internal"
	"matcher/pkg"
	"matcher/tools/zap"
	"strconv"
)

const driverServicePort = 8080

type IHandler interface {
	GetNearestDriver(c *fiber.Ctx) error
	ErrorHandler(c *fiber.Ctx, err error) error
}

type Handler struct {
	Repository internal.IRepository
}

func NewHandler(repository internal.IRepository) IHandler {
	return &Handler{
		Repository: repository,
	}
}

// @Summary Get nearest driver
// @Description Get nearest driver
// @Tags drivers
// @Accept json
// @Produce json
// @Param lat query string true "Latitude"
// @Param long query string true "Longitude"
// @Success 200 {object} pkg.Response
// @Failure 400 {object} pkg.Response
// @Failure 404 {object} pkg.Response
// @Failure 500 {object} pkg.Response
// @Router /drivers/nearest [get]
func (h *Handler) GetNearestDriver(c *fiber.Ctx) error {

	lat := c.Query("lat")
	long := c.Query("long")

	if lat == "" || long == "" {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.NewResponse(false, pkg.ErrInvalidRequest.Error(), nil))
	}

	floatLat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		zap.Logger.Error("Error parsing latitude: " + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(pkg.NewResponse(false, pkg.ErrInvalidRequest.Error(), nil))
	}

	floatLong, err := strconv.ParseFloat(long, 64)
	if err != nil {
		zap.Logger.Error("Error parsing longitude: " + err.Error())
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

		zap.Logger.Error("Error getting nearest driver: " + err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(pkg.NewResponse(false, pkg.ErrInternalServer.Error(), nil))
	}

	return c.Status(fiber.StatusOK).JSON(pkg.NewResponse(true, "Driver found", location))
}

func (h *Handler) ErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Log error for debugging
	zap.Logger.Error(err.Error())

	// Send error back
	return c.Status(code).JSON(pkg.NewResponse(false, pkg.ErrInternalServer.Error(), nil))
}

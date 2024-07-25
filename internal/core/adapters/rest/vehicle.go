package rest

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type vehicleHdl struct {
	vehicleService ports.VehicleService
}

func NewVehicleHdl(router *echo.Echo, vehicleService ports.VehicleService) ports.VehicleHandler {
	handler := &vehicleHdl{
		vehicleService: vehicleService,
	}

	api := router.Group("/vehicles")
	api.POST("/", handler.RegisterVehicle)
	api.GET("/", handler.ListSaleVehicles)
	api.GET("/sold", handler.ListSoldVehicles)
	api.PUT("/", handler.EditVehicle)

	return handler
}

func (h *vehicleHdl) RegisterVehicle(c echo.Context) error {
	var request VehicleRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})
	}

	id, err := h.vehicleService.RegisterVehicle(c.Request().Context(), domain.Vehicle{
		Make:         request.Make,
		Model:        request.Model,
		Year:         request.Year,
		Color:        request.Color,
		Price:        request.Price,
		Available:    true,
		CreationDate: time.Now(),
		ModifiedDate: time.Now(),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, GenericId{
		Id: id,
	})
}

func (h *vehicleHdl) ListSaleVehicles(c echo.Context) error {
	vehicles, err := h.vehicleService.ListSaleVehicles(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, vehicles)
}

func (h *vehicleHdl) ListSoldVehicles(c echo.Context) error {
	vehicles, err := h.vehicleService.ListSoldVehicles(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, vehicles)
}

func (h *vehicleHdl) EditVehicle(c echo.Context) error {
	var request VehicleRequest

	id := c.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})

	}

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})
	}

	err = h.vehicleService.EditVehicle(c.Request().Context(), domain.Vehicle{
		ID:        id,
		Make:      request.Make,
		Model:     request.Model,
		Year:      request.Year,
		Color:     request.Color,
		Price:     request.Price,
		Available: true,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.NoContent(http.StatusNoContent)
}

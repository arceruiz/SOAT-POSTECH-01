package rest

import (
	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/services"
	"github.com/labstack/echo/v4"
)

type VehicleHandler struct {
	vehicleService services.VehicleService
}

func NewVehicleHandler(g *echo.Group, vehicleService services.VehicleService) {
	handler := &VehicleHandler{
		vehicleService: vehicleService,
	}

	api := g.Group("/vehicles")
	api.POST("/", handler.CreateVehicle)
	api.GET("/", handler.GetAllVehicles)
	api.GET("/:id", handler.GetVehicle)
}

func (h *VehicleHandler) CreateVehicle(c echo.Context) error {
	//h.vehicleService.CreateVehicle()
	return nil
}

func (h *VehicleHandler) GetVehicle(c echo.Context) error {
	return nil
}

func (h *VehicleHandler) GetAllVehicles(c echo.Context) error {
	return nil
}

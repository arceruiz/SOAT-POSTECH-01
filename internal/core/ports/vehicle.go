package ports

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"context"

	"github.com/labstack/echo/v4"
)

type VehicleHandler interface {
	RegisterVehicle(echo.Context) error
	ListSaleVehicles(echo.Context) error
	ListSoldVehicles(echo.Context) error
	EditVehicle(echo.Context) error
}

type VehicleService interface {
	RegisterVehicle(context.Context, domain.Vehicle) (string, error)
	ListSaleVehicles(context.Context) ([]domain.Vehicle, error)
	ListSoldVehicles(context.Context) ([]domain.Vehicle, error)
	EditVehicle(context.Context, domain.Vehicle) error
}

type VehicleRepository interface {
	BuscarPorId(context.Context, string) (*domain.Vehicle, error)
	EditVehicle(context.Context, domain.Vehicle) error
	RegisterVehicle(ctx context.Context, vehicle domain.Vehicle) error
	ListSaleVehicles(context.Context) ([]domain.Vehicle, error)
	ListSoldVehicles(context.Context) ([]domain.Vehicle, error)
}

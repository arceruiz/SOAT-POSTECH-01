package ports

import "github.com/arceruiz/SOAT-POSTECH-01/internal/core/domain"

type VehicleService interface {
	CreateVehicle(vehicle domain.Vehicle) error
	GetVehicle(id string) (domain.Vehicle, error)
	GetAllVehicles() ([]domain.Vehicle, error)
}

type VehicleRepository interface {
	CreateVehicle(vehicle domain.Vehicle) error
	GetVehicle(id string) (domain.Vehicle, error)
	GetAllVehicles() ([]domain.Vehicle, error)
}

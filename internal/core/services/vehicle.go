package services

import (
	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/domain"
	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/ports"
)

type VehicleService struct {
	repo ports.VehicleRepository
}

func NewVehiclesService(repo ports.VehicleRepository) *VehicleService {
	return &VehicleService{repo: repo}
}

func (s *VehicleService) CreateVehicle(vehicle domain.Vehicle) error {
	return s.repo.CreateVehicle(vehicle)
}

func (s *VehicleService) GetVehicle(id string) (domain.Vehicle, error) {
	return s.repo.GetVehicle(id)
}

func (s *VehicleService) GetAllVehicles() ([]domain.Vehicle, error) {
	return s.repo.GetAllVehicles()
}

package usecases

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type vehicleService struct {
	repo ports.VehicleRepository
}

func NewVehicleService(repo ports.VehicleRepository) ports.VehicleService {
	return &vehicleService{repo: repo}
}

func (s *vehicleService) RegisterVehicle(ctx context.Context, vehicle domain.Vehicle) (string, error) {
	vehicle.ID = uuid.New().String()
	vehicle.CreationDate = time.Now()
	vehicle.ModifiedDate = time.Now()

	err := s.repo.RegisterVehicle(ctx, vehicle)
	if err != nil {
		logrus.WithError(err).Error("error saving vehicle")
		return "", err
	}

	return vehicle.ID, nil
}

func (s *vehicleService) ListSaleVehicles(ctx context.Context) ([]domain.Vehicle, error) {
	vehicles, err := s.repo.ListSaleVehicles(ctx)
	if err != nil {
		logrus.WithError(err).Error("error listing sale vehicles")
		return nil, err
	}

	return vehicles, nil
}

func (s *vehicleService) ListSoldVehicles(ctx context.Context) ([]domain.Vehicle, error) {
	vehicles, err := s.repo.ListSoldVehicles(ctx)
	if err != nil {
		logrus.WithError(err).Error("error listing sold vehicles")
		return nil, err
	}

	return vehicles, nil
}

func (s *vehicleService) EditVehicle(ctx context.Context, vehicle domain.Vehicle) error {
	veic, err := s.repo.BuscarPorId(ctx, vehicle.ID)
	if err != nil {
		logrus.WithError(err).Error("error searching vehicle")
		return err
	}

	veic.ModifiedDate = time.Now()

	err = s.repo.EditVehicle(ctx, vehicle)
	if err != nil {
		logrus.WithError(err).Error("error updating vehicle")
		return err
	}

	return nil
}

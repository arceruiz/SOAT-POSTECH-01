package usecases

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type saleService struct {
	repo              ports.SaleRepository
	vehicleRepository ports.VehicleRepository
}

func NewSaleService(repo ports.SaleRepository, vehicleRepo ports.VehicleRepository) ports.SaleService {
	return &saleService{repo: repo, vehicleRepository: vehicleRepo}
}

func (s *saleService) CreateSale(ctx context.Context, sale domain.Sale) error {
	sale.ID = uuid.New().String()
	sale.Date = time.Now()

	err := s.repo.CreateSale(ctx, sale)
	if err != nil {
		logrus.WithError(err).Error("error creating sale")
		return err
	}

	veic, err := s.vehicleRepository.BuscarPorId(ctx, sale.Vehicle.ID)
	if err != nil {
		logrus.WithError(err).Error("error searching vehicle")
		return err
	}

	veic.Available = false
	veic.ModifiedDate = time.Now()

	err = s.vehicleRepository.EditVehicle(ctx, *veic)
	if err != nil {
		logrus.WithError(err).Error("error editing vehicle")
		return err
	}

	return nil
}

func (s *saleService) SearchSales(ctx context.Context) ([]domain.Sale, error) {
	sales, err := s.repo.SearchSales(ctx)
	if err != nil {
		logrus.WithError(err).Error("error searching sales")
		return nil, err
	}

	return sales, nil
}

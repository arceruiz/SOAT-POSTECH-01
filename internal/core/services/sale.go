package services

import (
	"time"

	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/domain"
	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/ports"
	"github.com/google/uuid"
)

type SaleService struct {
	repo ports.SaleRepository
}

func NewSaleService(repo ports.SaleRepository) *SaleService {
	return &SaleService{repo: repo}
}

func (s *SaleService) CreateSale(sale domain.Sale) error {
	sale.ID = uuid.New().String()
	sale.Date = time.Now().Format("2006-01-02")
	return s.repo.CreateSale(sale)
}

func (s *SaleService) GetSale(id string) (domain.Sale, error) {
	return s.repo.GetSale(id)
}

func (s *SaleService) GetAllSales() ([]domain.Sale, error) {
	return s.repo.GetAllSales()
}

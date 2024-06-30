package ports

import "github.com/arceruiz/SOAT-POSTECH-01/internal/core/domain"

type SaleService interface {
	CreateSale(sale domain.Sale) error
	GetSale(id string) (domain.Sale, error)
	GetAllSales() ([]domain.Sale, error)
}

type SaleRepository interface {
	CreateSale(sale domain.Sale) error
	GetSale(id string) (domain.Sale, error)
	GetAllSales() ([]domain.Sale, error)
}

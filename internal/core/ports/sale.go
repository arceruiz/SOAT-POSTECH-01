package ports

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"context"

	"github.com/labstack/echo/v4"
)

type SaleHandler interface {
	SearchSales(echo.Context) error
	CreateSale(echo.Context) error
}
type SaleService interface {
	SearchSales(ctx context.Context) ([]domain.Sale, error)
	CreateSale(ctx context.Context, sale domain.Sale) error
}

type SaleRepository interface {
	SearchSales(ctx context.Context) ([]domain.Sale, error)
	CreateSale(ctx context.Context, sale domain.Sale) error
}

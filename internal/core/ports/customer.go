package ports

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"context"

	"github.com/labstack/echo/v4"
)

type CustomerHandler interface {
	RegisterCustomer(echo.Context) error
}

type CustomerService interface {
	RegisterCustomer(ctx context.Context, customer domain.Customer) (string, error)
}

type CustomerRepository interface {
	RegisterCustomer(ctx context.Context, customer domain.Customer) error
}

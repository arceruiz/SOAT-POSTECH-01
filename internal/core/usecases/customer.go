package usecases

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type customerService struct {
	repo ports.CustomerRepository
}

func NewCustomerService(repo ports.CustomerRepository) ports.CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) RegisterCustomer(ctx context.Context, customer domain.Customer) (string, error) {
	customer.ID = uuid.New().String()
	customer.CreationDate = time.Now()

	err := s.repo.RegisterCustomer(ctx, customer)
	if err != nil {
		logrus.WithError(err).Error("error saving customer")
		return "", err
	}

	return customer.ID, nil
}

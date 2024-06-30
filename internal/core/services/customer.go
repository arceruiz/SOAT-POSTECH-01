package services

import (
	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/domain"
	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/ports"
	"github.com/google/uuid"
)

type CustomerService struct {
	repo ports.CustomerRepository
}

func NewCustomerService(repo ports.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(customer domain.Customer) error {
	customer.ID = uuid.New().String()
	return s.repo.CreateCustomer(customer)
}

func (s *CustomerService) GetCustomer(id string) (domain.Customer, error) {

	return s.repo.GetCustomer(id)
}

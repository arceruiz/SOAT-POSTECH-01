package ports

import "github.com/arceruiz/SOAT-POSTECH-01/internal/core/domain"

type CustomerService interface {
	CreateCustomer(customer domain.Customer) error
	GetCustomer(id string) (domain.Customer, error)
}

type CustomerRepository interface {
	CreateCustomer(customer domain.Customer) error
	GetCustomer(id string) (domain.Customer, error)
}

package repositories

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type customerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) ports.CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) RegisterCustomer(ctx context.Context, customer domain.Customer) error {
	sqlStatement := "INSERT INTO \"Customer\" (id, name, document, creation_date) VALUES ($1, $2, $3, $4)"

	_, err := r.db.Exec(ctx, sqlStatement, customer.ID, customer.Name, customer.Document, customer.CreationDate)
	if err != nil {
		return err
	}

	return nil
}

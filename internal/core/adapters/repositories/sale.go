package repositories

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type saleRepository struct {
	db *pgxpool.Pool
}

func NewSaleRepository(db *pgxpool.Pool) ports.SaleRepository {
	return &saleRepository{db: db}
}

func (r *saleRepository) CreateSale(ctx context.Context, sale domain.Sale) error {
	sqlStatement := "INSERT INTO \"Sale\" (id, data, vehicle_id, customer_id, payment_type) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.db.Exec(ctx, sqlStatement, sale.ID, sale.Date, sale.Vehicle.ID, sale.Customer.ID, sale.PaymentType)
	if err != nil {
		return err
	}

	return nil
}

func (r *saleRepository) SearchSales(ctx context.Context) ([]domain.Sale, error) {
	rows, err := r.db.Query(ctx, "SELECT v.id, v.make, v.model, v.year, v.paint, v.price, v.available, c.id, c.name, c.document, c.creation_date, vd.id, vd.data, vd.payment_type FROM \"Sale\" vd INNER JOIN \"Vehicle\" v ON vd.vehicle_id = v.id INNER JOIN \"Customer\" c ON vd.customer_id = c.id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var sales []domain.Sale

	for rows.Next() {
		var sale domain.Sale

		if err = rows.Scan(
			&sale.Vehicle.ID,
			&sale.Vehicle.Make,
			&sale.Vehicle.Model,
			&sale.Vehicle.Year,
			&sale.Vehicle.Color,
			&sale.Vehicle.Price,
			&sale.Vehicle.Available,
			&sale.Customer.ID,
			&sale.Customer.Name,
			&sale.Customer.Document,
			&sale.Customer.CreationDate,
			&sale.ID,
			&sale.Date,
			&sale.PaymentType,
		); err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	return sales, nil
}

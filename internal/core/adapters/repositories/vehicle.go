package repositories

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type vehicleRepository struct {
	db *pgxpool.Pool
}

func NewVehicleRepository(db *pgxpool.Pool) ports.VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) RegisterVehicle(ctx context.Context, vehicle domain.Vehicle) error {
	sqlStatement := "INSERT INTO \"Vehicle\" (id, make, model, year, paint, price, available, creation_date, modified_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := r.db.Exec(ctx, sqlStatement, vehicle.ID, vehicle.Make, vehicle.Model, vehicle.Year, vehicle.Color, vehicle.Price, vehicle.Available, vehicle.CreationDate, vehicle.ModifiedDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *vehicleRepository) ListSaleVehicles(ctx context.Context) ([]domain.Vehicle, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM \"Vehicle\" WHERE available = true ORDER BY price ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var vehicles []domain.Vehicle

	for rows.Next() {
		var vehicle domain.Vehicle

		if err = rows.Scan(
			&vehicle.ID,
			&vehicle.Make,
			&vehicle.Model,
			&vehicle.Year,
			&vehicle.Color,
			&vehicle.Price,
			&vehicle.Available,
			&vehicle.CreationDate,
			&vehicle.ModifiedDate,
		); err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}

func (r *vehicleRepository) ListSoldVehicles(ctx context.Context) ([]domain.Vehicle, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM \"Vehicle\" WHERE available = false ORDER BY price ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var vehicles []domain.Vehicle

	for rows.Next() {
		var vehicle domain.Vehicle

		if err = rows.Scan(
			&vehicle.ID,
			&vehicle.Make,
			&vehicle.Model,
			&vehicle.Year,
			&vehicle.Color,
			&vehicle.Price,
			&vehicle.Available,
			&vehicle.CreationDate,
			&vehicle.ModifiedDate,
		); err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}

func (r *vehicleRepository) EditVehicle(ctx context.Context, vehicle domain.Vehicle) error {
	sqlStatement := "UPDATE \"Vehicle\" SET make = $1, model = $2, year = $3, paint = $4, price = $5, available = $6, creation_date = $7, modified_date = $8 WHERE id = $9"
	_, err := r.db.Exec(ctx, sqlStatement,
		vehicle.Make, vehicle.Model, vehicle.Year, vehicle.Color, vehicle.Price, vehicle.Available, vehicle.CreationDate, vehicle.ModifiedDate, vehicle.ID)
	if err != nil {
		return fmt.Errorf("erro ao atualizar o veículo: %w", err)
	}

	return nil
}

func (r *vehicleRepository) BuscarPorId(ctx context.Context, id string) (*domain.Vehicle, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Vehicle\" WHERE id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicle domain.Vehicle
	if rows.Next() {
		if err = rows.Scan(
			&vehicle.ID,
			&vehicle.Make,
			&vehicle.Model,
			&vehicle.Year,
			&vehicle.Color,
			&vehicle.Price,
			&vehicle.Available,
			&vehicle.CreationDate,
			&vehicle.ModifiedDate,
		); err != nil {
			return nil, err
		}
		return &vehicle, nil
	}

	return nil, err
}

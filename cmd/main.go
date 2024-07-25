package main

import (
	"SOAT-POSTECH-01/internal/config"
	"SOAT-POSTECH-01/internal/core/adapters/repositories"
	"SOAT-POSTECH-01/internal/core/adapters/rest"
	"SOAT-POSTECH-01/internal/core/ports"
	"SOAT-POSTECH-01/internal/core/usecases"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ParseFromFlags()

	router := echo.New()

	customerService, saleService, vehicleService := startDependencies()

	rest.NewCustomerHdl(router, customerService)
	rest.NewSaleHdl(router, saleService)
	rest.NewVehicleHdl(router, vehicleService)

	router.Start(":8081")
}

func startDependencies() (ports.CustomerService, ports.SaleService, ports.VehicleService) {
	db, err := pgxpool.New(context.Background(), config.Get().DB.ConnectionString)
	if err != nil {
		logrus.Fatal(err)
	}

	customerRepository := repositories.NewCustomerRepository(db)
	customerService := usecases.NewCustomerService(customerRepository)
	vehicleRepository := repositories.NewVehicleRepository(db)
	vehicleService := usecases.NewVehicleService(vehicleRepository)
	saleRepository := repositories.NewSaleRepository(db)
	saleService := usecases.NewSaleService(saleRepository, vehicleRepository)

	return customerService, saleService, vehicleService
}

package rest

import "time"

type CustomerRequest struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}

type GenericId struct {
	Id string `json:"id"`
}

type VehicleRequest struct {
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	Color        string    `json:"paint"`
	Price        float64   `json:"price"`
	CreationDate time.Time `json:"creation_date"`
	ModifiedDate time.Time `json:"modified_date"`
}

type SaleRequest struct {
	VehicleId   string `json:"vehicle_id"`
	CustomerId  string `json:"customer_id"`
	PaymentType string `json:"payment_type"`
}

package domain

import "time"

type Vehicle struct {
	ID           string
	Make         string
	Model        string
	Year         int
	Color        string
	Price        float64
	Available    bool
	CreationDate time.Time
	ModifiedDate time.Time
}

type Customer struct {
	ID           string
	Name         string
	Document     string
	CreationDate time.Time
}

type Sale struct {
	ID          string
	Date        time.Time
	Vehicle     Vehicle
	Customer    Customer
	PaymentType string
}

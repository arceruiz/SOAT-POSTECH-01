package domain

type Vehicle struct {
	ID          string
	Brand       string
	Model       string
	Year        int
	Color       string
	Price       float64
	IsAvailable bool
}

type Customer struct {
	ID       string
	Name     string
	Document string
}

type Sale struct {
	ID          string
	Date        string
	Vehicle     Vehicle
	Customer    Customer
	PaymentType string
}

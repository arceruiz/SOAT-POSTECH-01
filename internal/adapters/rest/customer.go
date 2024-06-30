package rest

import (
	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/services"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(g *echo.Group, customerService services.CustomerService) {
	handler := &CustomerHandler{
		customerService: customerService,
	}

	api := g.Group("/customers")
	api.POST("/", handler.CreateCustomer)
	api.GET("/", handler.GetAllCustomers)
	api.GET("/:id", handler.GetCustomer)
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	return nil
}

func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	return nil
}

func (h *CustomerHandler) GetAllCustomers(c echo.Context) error {
	return nil
}

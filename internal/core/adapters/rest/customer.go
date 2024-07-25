package rest

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type customerHdl struct {
	customerService ports.CustomerService
}

func NewCustomerHdl(router *echo.Echo, customerService ports.CustomerService) ports.CustomerHandler {
	handler := &customerHdl{
		customerService: customerService,
	}

	api := router.Group("/customers")
	api.POST("/", handler.RegisterCustomer)

	return handler
}

func (h *customerHdl) RegisterCustomer(c echo.Context) error {
	var request CustomerRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})
	}

	id, err := h.customerService.RegisterCustomer(c.Request().Context(), domain.Customer{
		Name:     request.Name,
		Document: request.Document,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, GenericId{
		Id: id,
	})
}

package rest

import (
	"SOAT-POSTECH-01/internal/core/domain"
	"SOAT-POSTECH-01/internal/core/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SaleHdl struct {
	saleService ports.SaleService
}

func NewSaleHdl(router *echo.Echo, SaleService ports.SaleService) ports.SaleHandler {
	handler := &SaleHdl{
		saleService: SaleService,
	}

	api := router.Group("/sale")
	api.POST("/", handler.CreateSale)
	api.GET("/", handler.SearchSales)

	return handler
}

func (h *SaleHdl) CreateSale(c echo.Context) error {
	var request SaleRequest

	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, struct{ string }{"invalid request"})
	}

	err = h.saleService.CreateSale(c.Request().Context(), domain.Sale{
		Vehicle: domain.Vehicle{
			ID: request.VehicleId,
		},
		Customer: domain.Customer{
			ID: request.CustomerId,
		},
		PaymentType: request.PaymentType,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.NoContent(http.StatusOK)
}

func (h *SaleHdl) SearchSales(c echo.Context) error {
	sales, err := h.saleService.SearchSales(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, struct{ string }{"unexpected error occurred"})
	}

	return c.JSON(http.StatusOK, sales)
}

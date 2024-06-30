package rest

import (
	"github.com/arceruiz/SOAT-POSTECH-01/internal/core/services"
	"github.com/labstack/echo/v4"
)

type SaleHandler struct {
	saleService services.SaleService
}

func NewSaleHandler(g *echo.Group, saleService services.SaleService) {
	handler := &SaleHandler{
		saleService: saleService,
	}

	api := g.Group("/sales")
	api.POST("/", handler.CreateSale)
	api.GET("/", handler.GetAllSales)
	api.GET("/:id", handler.GetSale)
}

func (h *SaleHandler) CreateSale(c echo.Context) error {
	return nil
}

func (h *SaleHandler) GetSale(c echo.Context) error {
	return nil
}

func (h *SaleHandler) GetAllSales(c echo.Context) error {
	return nil
}

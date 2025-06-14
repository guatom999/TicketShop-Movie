package inventoryHandlers

import (
	"context"
	"net/http"

	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryUseCases"
	"github.com/labstack/echo/v4"
)

type (
	InventoryHandlerService interface {
		GetCustomerTicket(c echo.Context) error
		FindLastCustomerTicket(c echo.Context) error
		HealthCheck(c echo.Context) error
	}

	inventoryHandler struct {
		inventoryUseCase inventoryUseCases.InventoryUseCaseService
	}
)

func NewInventoryHandler(inventoryUseCase inventoryUseCases.InventoryUseCaseService) InventoryHandlerService {
	return &inventoryHandler{inventoryUseCase: inventoryUseCase}
}

func (h *inventoryHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "status ok")
}

func (h *inventoryHandler) GetCustomerTicket(c echo.Context) error {

	ctx := context.Background()

	customerId := c.Param("customerid")

	results, err := h.inventoryUseCase.GetCustomerTicket(ctx, customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

func (h *inventoryHandler) GetCustomerTicketDetail(c echo.Context) error {

	return nil

}

func (h *inventoryHandler) FindLastCustomerTicket(c echo.Context) error {

	ctx := context.Background()

	customerId := c.Param("customerid")

	results, err := h.inventoryUseCase.FindLastCustomerTicket(ctx, customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

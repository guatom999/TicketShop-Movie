package inventoryHandlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/guatom999/TicketShop-Movie/modules/inventory/inventoryUseCases"
	"github.com/labstack/echo/v4"
)

type (
	InventoryHandlerService interface {
		FindCustomerTicket(c echo.Context) error
	}

	inventoryHandler struct {
		inventoryUseCase inventoryUseCases.InventoryUseCaseService
	}
)

func NewInventoryHandler(inventoryUseCase inventoryUseCases.InventoryUseCaseService) InventoryHandlerService {
	return &inventoryHandler{inventoryUseCase: inventoryUseCase}
}

func (h *inventoryHandler) FindCustomerTicket(c echo.Context) error {

	ctx := context.Background()

	customerId := c.Param("customerid")

	fmt.Println("FindCustomerTicket is +====================================>", customerId)

	_, err := h.inventoryUseCase.FindCustomerTicket(ctx, customerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "test success")
}

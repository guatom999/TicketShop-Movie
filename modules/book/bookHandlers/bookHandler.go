package bookHandlers

import (
	"net/http"

	"github.com/guatom999/TicketShop-Movie/modules/book/bookUseCases"
	"github.com/labstack/echo/v4"
)

type (
	BookHandlerService interface {
		BuyTicket(c echo.Context) error
	}

	bookHandler struct {
		bookUseCases bookUseCases.BookUseCaseService
	}
)

func NewBookHandler(bookUseCases bookUseCases.BookUseCaseService) BookHandlerService {
	return &bookHandler{bookUseCases: bookUseCases}
}

func (u *bookHandler) BuyTicket(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

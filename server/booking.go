package server

import (
	"net/http"

	"github.com/guatom999/TicketShop-Movie/modules/book/bookHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/book/bookRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/book/bookUseCases"
	"github.com/labstack/echo/v4"
)

func (s *server) BookingModule() {
	bookingRepo := bookRepositories.NewBookRepository()
	bookingUseCase := bookUseCases.NewBookUsecase(bookingRepo)
	bookingHandler := bookHandlers.NewBookHandler(bookingUseCase)

	_ = bookingHandler

	bookingRouter := s.app.Group("/booking")

	bookingRouter.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Test Success")
	})

}

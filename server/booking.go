package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/book/bookHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/book/bookRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/book/bookUseCases"
	"github.com/guatom999/TicketShop-Movie/pkg/queue"
	"github.com/labstack/echo/v4"
)

func (s *server) BookingModule() {
	bookingRepo := bookRepositories.NewBookRepository()
	bookingUseCase := bookUseCases.NewBookUsecase(bookingRepo)
	bookingHandler := bookHandlers.NewBookHandler(bookingUseCase)

	_ = bookingHandler

	bookingConsumer(s.cfg)

	bookingRouter := s.app.Group("/booking")

	bookingRouter.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Test Success")
	})

}

func bookingConsumer(cfg *config.Config) {

	conn := queue.KafkaConn(cfg)

	for {
		message, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Println(string(message.Value))
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}

}

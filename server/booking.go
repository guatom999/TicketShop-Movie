package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/book/bookHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/book/bookRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/book/bookUseCases"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
)

func (s *server) BookingModule() {
	bookingRepo := bookRepositories.NewBookRepository()
	bookingUseCase := bookUseCases.NewBookUsecase(bookingRepo)
	bookingHandler := bookHandlers.NewBookHandler(bookingUseCase)

	_ = bookingHandler

	go bookingConsumer(s.cfg)

	bookingRouter := s.app.Group("/booking")

	bookingRouter.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Test Success")
	})

}

func bookingConsumer(cfg *config.Config) {

	// conn := queue.KafkaConn(cfg)

	// offset, _ := conn

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.Kafka.Url},
		Topic:     "kafkaapikey",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	// conn

	fmt.Println("Reader is ", reader.SetOffset(6))

	for {

		// fmt.Printf("First Offset is %d lastOffset is %d ", fisrtOffset, lastOffset)
		// fmt.Println("Offset is", offset)
		// message, err := conn.ReadMessage(10e3)
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Println("Message is", string(message.Value))
	}

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close connection:", err)
	// }

}

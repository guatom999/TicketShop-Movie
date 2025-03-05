package moviesHandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesUseCases"
	"github.com/guatom999/TicketShop-Movie/pkg/queue"
)

type (
	MoviesQueueHandlerService interface {
		ReserveSeat()
		RollBackSeat()
	}

	moviesQueueHandler struct {
		cfg          *config.Config
		movieUseCase moviesUseCases.MoviesUseCaseService
	}
)

func NewMoviesQueueHandler(cfg *config.Config, movieUseCase moviesUseCases.MoviesUseCaseService) MoviesQueueHandlerService {
	return &moviesQueueHandler{
		cfg:          cfg,
		movieUseCase: movieUseCase,
	}
}

// func (h *moviesQueueHandler) MovieConsumer(pctx context.Context) *kafka.Reader {

// 	reader := queue.KafkaReader()

// 	return reader

// }

// func getConsumerMessage(consumer *kafka.Conn, consumerMessage chan kafka.Message) {

// 	message, err := consumer.ReadMessage(10e3)
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println("message from consumer is", message)

// 	consumerMessage <- message

// }

func (h *moviesQueueHandler) ReserveSeat() {

	ctx := context.Background()

	data := new(movie.ReserveSeatReqTest)

	fmt.Println("h.cfg", h.cfg.Kafka.ApiKey)

	reader := queue.KafkaReader(h.cfg, "reserve-seat", "reserve-seat-req-group")
	defer reader.Close()

	for {

		message, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %s", err.Error())
			break
		}

		if string(message.Key) == "movie" {

			if err := json.Unmarshal(message.Value, data); err != nil {
				fmt.Printf("Error: Unmarshal error %s", err.Error())
			}

			h.movieUseCase.ReserveSeat(ctx, &movie.ReserveDetailReq{
				MovieId: data.MovieId,
				SeatNo:  data.Seat_Number,
			})
		}
	}
}

func (h *moviesQueueHandler) RollBackSeat() {

	ctx := context.Background()

	data := new(movie.RollBackReservedSeatReq)

	reader := queue.KafkaReader(h.cfg, "rollback", "roll-back-seat-req-group")
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading message: %s", err.Error())
			break
		}

		if string(message.Key) == "movie" {
			if err := json.Unmarshal(message.Value, data); err != nil {
				fmt.Printf("Error: Unmarshal error %s", err.Error())
			}

			h.movieUseCase.RollbackReserveSeat(ctx, &movie.ReserveDetailReq{
				MovieId: data.MovieId,
				SeatNo:  data.SeatNo,
			})
		}
	}
}

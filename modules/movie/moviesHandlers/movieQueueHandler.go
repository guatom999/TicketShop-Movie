package moviesHandlers

import (
	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesUseCases"
)

type (
	MoviesQueueHandlerService interface {
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

// func MovieConsumer() {

// 	ctx := context.Background()

// 	conn := queue.KafkaConn()
// }

func (h *moviesQueueHandler) ReserveSeat() {

}

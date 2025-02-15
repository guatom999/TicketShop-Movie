package server

import (
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesUseCases"
)

func (s *server) MovieModule() {
	movieRepo := moviesRepositories.NewMoviesrepository(s.db, s.redis)
	movieUseCase := moviesUseCases.NewmoviesUseCase(s.cfg, movieRepo)
	movieHandler := moviesHandlers.NewMoviesHandler(movieUseCase)
	movieQueueHandler := moviesHandlers.NewMoviesQueueHandler(s.cfg, movieUseCase)

	go movieQueueHandler.ReserveSeat()
	go movieQueueHandler.RollBackSeat()

	movieRouter := s.app.Group("/movie")

	movieRouter.POST("/addmovie", movieHandler.AddOneMovie)
	movieRouter.GET("/getallmovie", movieHandler.GetAllMovie)
	movieRouter.GET("/comingsoonmovie", movieHandler.GetAllComingSoonMovie)
	movieRouter.GET("/getmovie/:movie_id", movieHandler.FindOneMovie)
	movieRouter.GET("/getmovieshowtime/:movieid", movieHandler.FindMovieShowTime)

	movieRouter.POST("/reserveseat", movieHandler.ReserveSeat)

	movieRouter.GET("/test", movieHandler.TestReq)

}

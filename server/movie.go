package server

import (
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesUseCases"
)

func (s *server) MovieModule() {
	movieRepo := moviesRepositories.NewMoviesrepository(s.db)
	movieUseCase := moviesUseCases.NewmoviesUseCase(movieRepo)
	movieHandler := moviesHandlers.NewMoviesHandler(movieUseCase)

	movieRouter := s.app.Group("/movie")

	movieRouter.POST("/addmovie", movieHandler.AddOneMovie)
	movieRouter.GET("/getallmovie", movieHandler.GetAllMovie)
	movieRouter.GET("/getmovie/:title", movieHandler.FindOneMovie)
	movieRouter.GET("/test", movieHandler.TestReq)
	movieRouter.GET("/getmovieshowtime/:title", movieHandler.FindMovieShowTime)

	//Test
	movieRouter.POST("/testreserveseat", movieHandler.ReserveSeat)
}
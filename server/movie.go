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

	movieRouter.POST("/add", movieHandler.AddOneMovie)
	movieRouter.GET("/getallmovie", movieHandler.GetAllMovie)
}

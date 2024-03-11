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
	movieRouter.GET("/getmovieshowtime/:movieid", movieHandler.FindMovieShowTime)

	//Test
	movieRouter.POST("/reserveseat", movieHandler.ReserveSeat)

	movieRouter.GET("/test", movieHandler.TestReq)

}

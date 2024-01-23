package moviesHandlers

import (
	"context"
	"net/http"

	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesUseCases"
	"github.com/labstack/echo/v4"
)

type (
	MoviesHandlerService interface {
		AddOneMovie(c echo.Context) error
	}

	moviesHandler struct {
		moviesUseCase moviesUseCases.MoviesUseCaseService
	}
)

func NewMoviesHandler(moviesUseCase moviesUseCases.MoviesUseCaseService) MoviesHandlerService {
	return &moviesHandler{moviesUseCase: moviesUseCase}
}

func (h *moviesHandler) AddOneMovie(c echo.Context) error {
	ctx := context.Background()

	req := new(movie.AddMovieReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.moviesUseCase.AddOneMovie(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "add movie success")

}

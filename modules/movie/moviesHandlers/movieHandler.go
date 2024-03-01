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
		FindOneMovie(c echo.Context) error
		GetAllMovie(c echo.Context) error
		TestReq(c echo.Context) error
		FindMovieShowTime(c echo.Context) error
		ReserveSeat(c echo.Context) error
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

func (h *moviesHandler) FindOneMovie(c echo.Context) error {

	ctx := context.Background()

	movieName := c.Param("title")

	movie, err := h.moviesUseCase.FindOneMovie(ctx, movieName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, movie)

}

func (h *moviesHandler) GetAllMovie(c echo.Context) error {

	ctx := context.Background()

	result, err := h.moviesUseCase.FindAllMovie(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *moviesHandler) FindMovieShowTime(c echo.Context) error {

	ctx := context.Background()

	movieName := c.Param("title")

	movies, err := h.moviesUseCase.FindMovieShowTime(ctx, movieName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, movies)
}

func (h *moviesHandler) ReserveSeat(c echo.Context) error {

	ctx := context.Background()

	input := make([]*movie.ReserveDetailReq, 0)

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.moviesUseCase.ReserveSeat(ctx, input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "test success")
}

func (h *moviesHandler) TestReq(c echo.Context) error {

	ctx := context.Background()

	result, err := h.moviesUseCase.TestReq(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)
}

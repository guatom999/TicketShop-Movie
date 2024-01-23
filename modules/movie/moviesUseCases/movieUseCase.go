package moviesUseCases

import (
	"context"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesRepositories"
)

type (
	MoviesUseCaseService interface {
		AddOneMovie(pctx context.Context, req *movie.AddMovieReq) error
	}

	moviesUseCase struct {
		moviesRepo moviesRepositories.MoviesRepositoryService
	}
)

func NewmoviesUseCase(moviesRepo moviesRepositories.MoviesRepositoryService) MoviesUseCaseService {
	return &moviesUseCase{moviesRepo: moviesRepo}
}

func (u *moviesUseCase) AddOneMovie(pctx context.Context, req *movie.AddMovieReq) error {

	if err := u.moviesRepo.InsertMovie(pctx, &movie.Movie{
		Title:     req.Title,
		Price:     req.Price,
		ImageUrl:  req.ImageUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Avaliable: req.Avaliable,
	}); err != nil {
		return err
	}

	return nil
}

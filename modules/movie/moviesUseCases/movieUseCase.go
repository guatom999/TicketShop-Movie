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
		FindAllMovie(pctx context.Context) ([]*movie.MovieData, error)
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

func (u *moviesUseCase) FindAllMovie(pctx context.Context) ([]*movie.MovieData, error) {

	result, err := u.moviesRepo.FindAllMovie(pctx)
	if err != nil {
		return make([]*movie.MovieData, 0), nil
	}

	return result, nil
}

func (u *moviesUseCase) FindManyMovie(pctx context.Context, basePaginateUrl string) error {
	return nil
}

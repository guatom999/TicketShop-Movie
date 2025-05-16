package moviesRepositories

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/stretchr/testify/mock"
)

type (
	MovieRepositoryMock struct {
		mock.Mock
	}
)

func NewMovieRepoMock() MoviesRepositoryService {
	return &MovieRepositoryMock{}
}

func (m *MovieRepositoryMock) InsertMovie(pctx context.Context, req []*movie.Movie, movieRound int64) error {
	args := m.Called(pctx, req, movieRound)
	return args.Error(0)
}
func (m *MovieRepositoryMock) FindOneMovie(pctx context.Context, movieId string) (*movie.MovieShowCase, error) {
	args := m.Called(pctx, movieId)
	return args.Get(0).(*movie.MovieShowCase), args.Error(1)
}
func (m *MovieRepositoryMock) FindAllMovie(pctx context.Context, filter any) ([]*movie.MovieData, error) {
	args := m.Called(pctx, filter)
	return args.Get(0).([]*movie.MovieData), args.Error(1)
}
func (m *MovieRepositoryMock) FindComingSoonMovie(pctx context.Context, filter any) ([]*movie.MovieData, error) {
	args := m.Called(pctx, filter)
	return args.Get(0).([]*movie.MovieData), args.Error(1)
}
func (m *MovieRepositoryMock) FindMovieShowtime(pctx context.Context, movieId string) ([]*movie.MovieShowTimeRes, error) {
	args := m.Called(pctx, movieId)
	return args.Get(0).([]*movie.MovieShowTimeRes), args.Error(1)
}
func (m *MovieRepositoryMock) GetOneMovieAvaliable(pctx context.Context, req *movie.ReserveDetailReq) (*movie.MovieAvaliable, error) {
	args := m.Called(pctx, req)
	return args.Get(0).(*movie.MovieAvaliable), args.Error(1)
}
func (m *MovieRepositoryMock) UpdateSeatStatus(pctx context.Context, movidId string, req *movie.MovieAvaliable) error {
	args := m.Called(pctx, movidId, req)
	return args.Error(0)
}
func (m *MovieRepositoryMock) ReserveSeatRes(pctx context.Context, cfg *config.Config, req *movie.ReserveSeatRes) error {
	args := m.Called(pctx, cfg, req)
	return args.Error(0)
}

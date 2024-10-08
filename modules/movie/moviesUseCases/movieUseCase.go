package moviesUseCases

import (
	"context"
	"fmt"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/guatom999/TicketShop-Movie/modules/movie/moviesRepositories"
	"github.com/guatom999/TicketShop-Movie/pkg/rest"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	MoviesUseCaseService interface {
		AddOneMovie(pctx context.Context, req *movie.AddMovieReq) error
		FindAllMovie(pctx context.Context) ([]*movie.MovieData, error)
		FindComingSoonMovie(pctx context.Context) ([]*movie.MovieData, error)
		TestReq(pctx context.Context) (string, error)
		FindOneMovie(pctx context.Context, movieId string) (*movie.MovieShowCase, error)
		FindMovieShowTime(pctx context.Context, title string) ([]*movie.MovieShowTimeRes, error)
		ReserveSeat(pctx context.Context, req *movie.ReserveDetailReq) error
	}

	moviesUseCase struct {
		moviesRepo moviesRepositories.MoviesRepositoryService
	}
)

func NewmoviesUseCase(moviesRepo moviesRepositories.MoviesRepositoryService) MoviesUseCaseService {
	return &moviesUseCase{moviesRepo: moviesRepo}
}

func (u *moviesUseCase) AddOneMovie(pctx context.Context, req *movie.AddMovieReq) error {

	// splitDate := strings.Split(req.ReleaseAt, "-")

	// for _, v := range splitDate {
	// 	fmt.Println("test splitdate", v)
	// }

	// fmt.Println("convert time is", utils.ConvertStringDateToTime(req.ReleaseAt))

	if err := u.moviesRepo.InsertMovie(pctx, &movie.Movie{
		Title:           req.Title,
		Description:     req.Description,
		RunningTime:     req.RunningTime,
		Price:           req.Price,
		ImageUrl:        req.ImageUrl,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Category:        "RomCom",
		ReleaseAt:       utils.ConvertStringDateToTime(req.ReleaseAt),
		OutOfTheatersAt: utils.ConvertStringDateToTime(req.OutOfTheatersAt),
	}); err != nil {
		return err
	}

	return nil
}

// FindWithMoreCondition
func (u *moviesUseCase) FindOneMovie(pctx context.Context, movieId string) (*movie.MovieShowCase, error) {

	result, err := u.moviesRepo.FindOneMovie(pctx, movieId)
	if err != nil {
		return nil, err
	}

	// return &movie.MovieShowCase{
	// 	Title:    result.Title,
	// 	Price:    result.Price,
	// 	ImageUrl: result.ImageUrl,
	// }, nil

	return result, nil

}

func (u *moviesUseCase) FindAllMovie(pctx context.Context) ([]*movie.MovieData, error) {

	filter := bson.D{}

	filter = append(filter, bson.E{"out_of_theaters_at", bson.D{{"$gt", utils.GetLocaltime()}}}, bson.E{"release_at", bson.D{{"$lt", utils.GetLocaltime()}}})

	result, err := u.moviesRepo.FindAllMovie(pctx, filter)
	if err != nil {
		return make([]*movie.MovieData, 0), nil
	}

	return result, nil
}

func (u *moviesUseCase) FindComingSoonMovie(pctx context.Context) ([]*movie.MovieData, error) {

	filter := bson.D{}

	filter = append(filter, bson.E{"release_at", bson.D{{"$gt", utils.GetLocaltime()}}})

	results, err := u.moviesRepo.FindComingSoonMovie(pctx, filter)
	if err != nil {
		return results, err
	}

	return results, nil
}

func (u *moviesUseCase) FindManyMovie(pctx context.Context, basePaginateUrl string) error {

	// findItemsFilter := bson.D{}
	// findItemOption := make([]*options.FindOptions, 0)

	// countItemsFilter := bson.D{}

	// // Find many item filter
	// if req.Start != "" {
	// 	req.Start = strings.TrimPrefix(req.Start, "item:")
	// 	findItemsFilter = append(findItemsFilter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	// }

	// if req.Title != "" {
	// 	findItemsFilter = append(findItemsFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
	// 	countItemsFilter = append(countItemsFilter, bson.E{"title", primitive.Regex{Pattern: req.Title, Options: "i"}})
	// }

	// findItemsFilter = append(findItemsFilter, bson.E{"usage_status", true})
	// countItemsFilter = append(countItemsFilter, bson.E{"usage_status", true})

	// // options
	// findItemOption = append(findItemOption, options.Find().SetSort(bson.D{{"_id", 1}}))
	// findItemOption = append(findItemOption, options.Find().SetLimit(int64(req.Limit)))

	// results, err := u.itemRepo.FindManyItems(pctx, findItemsFilter, findItemOption)
	// if err != nil {
	// 	return nil, err
	// }

	// total, err := u.itemRepo.CountItems(pctx, countItemsFilter)
	// if err != nil {
	// 	return nil, err
	// }

	// if len(results) == 0 {
	// 	return &models.PaginateRes{
	// 		Data:  make([]*item.ItemShowCase, 0),
	// 		Total: total,
	// 		Limit: req.Limit,
	// 		First: models.FirstPaginate{
	// 			Href: fmt.Sprintf("%s?limit=%d&title=%s", req.Limit, basePaginateUrl, req.Title),
	// 		},
	// 		Next: models.NextPaginate{
	// 			Start: "",
	// 			Href:  "",
	// 		},
	// 	}, nil
	// }

	// return &models.PaginateRes{
	// 	Data:  results,
	// 	Total: total,
	// 	Limit: req.Limit,
	// 	First: models.FirstPaginate{
	// 		Href: fmt.Sprintf("%s?limit=%d&title=%s", req.Limit, basePaginateUrl, req.Title),
	// 	},
	// 	Next: models.NextPaginate{
	// 		Start: results[len(results)-1].ItemId,
	// 		Href:  fmt.Sprintf("%s?limit=%d&title=%s&start=%s", basePaginateUrl, req.Limit, req.Title, results[len(results)-1].ItemId),
	// 	},
	// }, nil

	return nil
}

func (u *moviesUseCase) FindMovieShowTime(pctx context.Context, title string) ([]*movie.MovieShowTimeRes, error) {

	movies, err := u.moviesRepo.FindMovieShowtime(pctx, title)
	if err != nil {
		return make([]*movie.MovieShowTimeRes, 0), err
	}

	return movies, nil
}

func (u *moviesUseCase) ReserveSeat(pctx context.Context, req *movie.ReserveDetailReq) error {

	if err := u.moviesRepo.UpdateSeatStatus(pctx, req); err != nil {
		return err
	}

	return nil
}

func (u *moviesUseCase) TestReq(pctx context.Context) (string, error) {

	url := "http://localhost:8099/booking/test"

	res, err := rest.Request(url)
	if err != nil {

		fmt.Println("Error: Error is ", res)
		return "", err
	}

	fmt.Println("Response body:", res)

	return res, nil
}

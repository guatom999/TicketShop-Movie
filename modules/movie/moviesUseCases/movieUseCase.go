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

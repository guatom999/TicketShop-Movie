package moviesRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MoviesRepositoryService interface {
		InsertMovie(pctx context.Context, req *movie.Movie) error
		FindOneMovie(pctx context.Context, title string) (*movie.Movie, error)
		FindAllMovie(pctx context.Context, filter any) ([]*movie.MovieData, error)
		FindMovieShowtime(pctx context.Context, title string) ([]*movie.MovieShowTimeRes, error)
		// IsMovieAvaliable(pctx context.Context, req string) bool
	}

	moviesrepository struct {
		db *mongo.Client
	}
)

func NewMoviesrepository(db *mongo.Client) MoviesRepositoryService {
	return &moviesrepository{db: db}
}

func (r *moviesrepository) InsertMovie(pctx context.Context, req *movie.Movie) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: Insert One Movie Failed")
		return errors.New("error: insert one movie failed")
	}

	log.Println("item id is", result.InsertedID)

	return nil

}

// func (r *moviesrepository) IsMovieAvaliable(pctx context.Context, req string) bool {

// 	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
// 	defer cancel()

// 	db := r.db.Database("movie_db")
// 	col := db.Collection("movie")

// 	movie := new(movie.Movie)

// 	if err := col.FindOne(ctx, bson.M{"Title": req}).Decode(movie); err != nil {
// 		log.Printf("Error: IsMovieAvaliable Falied:%s", err.Error())
// 		return false
// 	}

// 	if len(movie.Movie) {
// 		return true
// 	}

// 	if movie.Avaliable <= 0 {
// 		log.Printf("Movie %s is out of stock", movie.Title)
// 		return false
// 	}

// 	return true
// }

func (r *moviesrepository) FindOneMovie(pctx context.Context, title string) (*movie.Movie, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	result := new(movie.Movie)

	if err := col.FindOne(ctx, bson.M{"Title": title}).Decode(result); err != nil {
		log.Printf("Error: FindOne Movie Failed:%s", err.Error())
		return nil, errors.New("error: findone movie failed")
	}

	return result, nil
}

func (r *moviesrepository) FindAllMovie(pctx context.Context, filter any) ([]*movie.MovieData, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		log.Printf("Error: Find All Movie Failed: %s", err.Error())
		return make([]*movie.MovieData, 0), errors.New("error: find all item failed")
	}

	results := make([]*movie.MovieData, 0)

	log.Println("result is", results)

	for cursor.Next(ctx) {
		result := new(movie.Movie)
		if err := cursor.Decode(result); err != nil {
			log.Println("Error: Find All Movie Failed:", err.Error())
			return make([]*movie.MovieData, 0), errors.New("error: find all item failed")
		}

		results = append(results, &movie.MovieData{
			MovieId:  result.Id.Hex(),
			Title:    result.Title,
			Price:    result.Price,
			ImageUrl: result.ImageUrl,
		})
	}

	return results, nil
}

func (r *moviesrepository) FindMovieShowtime(pctx context.Context, title string) ([]*movie.MovieShowTimeRes, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie_available")

	results := make([]*movie.MovieShowTimeRes, 0)

	log.Println(utils.GetLocaltime())

	cursor, err := col.Find(ctx, bson.M{"title": title})
	if err != nil {
		log.Printf("Error: FindOne Movie Failed:%s", err.Error())
		return nil, errors.New("error: findone movie failed")
	}

	for cursor.Next(ctx) {
		result := new(movie.MovieAvaliable)
		if err := cursor.Decode(result); err != nil {
			log.Printf("Error: Decode FineMovieShowtime failed:%s", err)
			return make([]*movie.MovieShowTimeRes, 0), errors.New("error: find movie showtime failed")
		}

		results = append(results, &movie.MovieShowTimeRes{
			Title:    result.Title,
			ShowTime: utils.GetStringTime(result.Showtime),
		})

	}

	return results, nil
}

// func (r *moviesrepository) FindManyMovies(pctx context.Context, filter primitive.ObjectID, opts []*options.FindOptions) ([]*movie.MovieData, error) {

// 	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
// 	defer cancel()

// 	db := r.db.Database("movie_db")
// 	col := db.Collection("movie")

// 	cur, err := col.Find(ctx, filter, opts...)
// 	if err != nil {
// 		log.Printf("Error: Find All Movie Failed: %s", err.Error())
// 		return make([]*movie.MovieData, 0), errors.New("error: find many item failed")
// 	}

// 	results := make([]*movie.MovieData, 0)

// 	for cur.Next(ctx) {
// 		result := new(movie.Movie)
// 		if err := cur.Decode(result); err != nil {
// 			log.Printf("Error: Find Many Movies Failed:%s", err.Error())
// 			return make([]*movie.MovieData, 0), errors.New("error: find many movie failed")
// 		}

// 		results = append(results, &movie.MovieData{
// 			MovieId:   result.Id.Hex(),
// 			Title:     result.Title,
// 			Price:     result.Price,
// 			ImageUrl:  result.ImageUrl,
// 			Avaliable: result.Avaliable,
// 		})

// 	}

// 	return results, nil
// }

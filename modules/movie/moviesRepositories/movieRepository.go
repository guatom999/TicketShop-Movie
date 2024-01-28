package moviesRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	MoviesRepositoryService interface {
		InsertMovie(pctx context.Context, req *movie.Movie) error
		FindAllMovie(pctx context.Context) ([]*movie.MovieData, error)
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

func (r *moviesrepository) FindAllMovie(pctx context.Context,
) ([]*movie.MovieData, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	filter := bson.D{}

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
			MovieId:   result.Id.Hex(),
			Title:     result.Title,
			Price:     result.Price,
			ImageUrl:  result.ImageUrl,
			Avaliable: result.Avaliable,
		})
	}

	return results, nil
}

func (r *moviesrepository) FindManyMovies(pctx context.Context, filter primitive.ObjectID, opts []*options.FindOptions) ([]*movie.MovieData, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	cur, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Printf("Error: Find All Movie Failed: %s", err.Error())
		return make([]*movie.MovieData, 0), errors.New("error: find many item failed")
	}

	results := make([]*movie.MovieData, 0)

	for cur.Next(ctx) {
		result := new(movie.Movie)
		if err := cur.Decode(result); err != nil {
			log.Printf("Error: Find Many Movies Failed:%s", err.Error())
			return make([]*movie.MovieData, 0), errors.New("error: find many movie failed")
		}

		results = append(results, &movie.MovieData{
			MovieId:   result.Id.Hex(),
			Title:     result.Title,
			Price:     result.Price,
			ImageUrl:  result.ImageUrl,
			Avaliable: result.Avaliable,
		})

	}

	return results, nil
}

package moviesRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MoviesRepositoryService interface {
		InsertMovie(pctx context.Context, req *movie.Movie) error
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

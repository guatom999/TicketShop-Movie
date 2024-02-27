package moviesRepositories

import (
	"context"
	"errors"
	"fmt"
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
		UpdateSeatStatus(pctx context.Context, req *movie.ReserveDetailReq) error
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
			Title:         result.Title,
			ShowTime:      utils.GetStringTime(result.Showtime),
			SeatAvailable: result.SeatAvailable,
		})

	}

	return results, nil
}

func (r *moviesrepository) UpdateSeatStatus(pctx context.Context, req *movie.ReserveDetailReq) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie_available")

	result := new(movie.MovieAvaliable)

	if err := col.FindOne(ctx, bson.M{"movie_id": req.MovieId}).Decode(result); err != nil {
		log.Printf("Error: Find Seat Status Failed:%s", err.Error())
		return errors.New("error: find seat status failed")
	}
	// req.SeatNo = "Z3"

	for i, seat := range result.SeatAvailable {

		if _, ok := seat[fmt.Sprint(req.SeatNo)]; ok {
			if seat[req.SeatNo] {
				log.Println("Update seat now")
				result.SeatAvailable[i][req.SeatNo] = false
				break
			}
		} else if i == (len(result.SeatAvailable) - 1) {
			log.Println("error:no seat match")
			return errors.New("error: no seat match")
		}
	}

	updateResult, err := col.UpdateOne(ctx, bson.M{"movie_id": req.MovieId}, bson.M{"$set": result})
	if err != nil {
		log.Printf("Error: Update Seat Status Failed %v", err)
		return errors.New("error: update seat status failed")
	}

	log.Printf("update status is :%v", updateResult)

	return nil

}

func (r *moviesrepository) ReserveSeat(pctx context.Context, req *movie.ReserveDetailReq) error {

	// ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	// defer cancel()

	// db := r.db.Database("movie_db")
	// col := db.Collection("movie_available")

	// if !r.IsSeatAvailable(ctx, &movie.ReserveDetailReq{}) {
	// 	log.Printf("Error: Seat is Not Available")
	// 	return errors.New("errro: seat is not available")
	// }

	// col.UpdateOne(ctx, bson.M{"movie_id": req.MovieId}, bson.D{{"$set", bson.D{{"seat_available"}}}})

	// col.UpdateOne(
	// 	ctx,
	// 	bson.M{"_id": utils.ConvertToObjectId(req.MovieId)},
	// 	bson.M{"$set": bson.M{
	// 		"seat_available":
	// 	}})

	return nil
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

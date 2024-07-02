package moviesRepositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MoviesRepositoryService interface {
		InsertMovie(pctx context.Context, req *movie.Movie) error
		FindOneMovie(pctx context.Context, movieId string) (*movie.Movie, error)
		FindAllMovie(pctx context.Context, filter any) ([]*movie.MovieData, error)
		FindComingSoonMovie(pctx context.Context, filter any) ([]*movie.MovieData, error)
		FindMovieShowtime(pctx context.Context, movieId string) ([]*movie.MovieShowTimeRes, error)
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

	lastInsertId := result.InsertedID.(primitive.ObjectID)

	col = db.Collection("movie_available")

	// movieAvailableCount, err := col.CountDocuments(pctx, bson.M{})
	// if err != nil {
	// 	log.Printf("Error: Count Documents failed: %v", err)
	// 	return errors.New("error: countDocuments failed")
	// }

	addMovieAvailable := func() []any {

		datas := make([]any, 0)

		for i := 0; i < 10; i++ {

			data := movie.MovieAvaliable{
				Movie_Id:  lastInsertId.Hex(),
				Title:     req.Title,
				CreatedAt: utils.GetLocaltime(),
				UpdatedAt: func() time.Time {
					return utils.GetLocaltime()
				}(),
				Showtime: utils.SetSpecificTime(2024, 3, 2+int(math.Floor(float64(i)/2)), 10+i, 30, 0),
				SeatAvailable: []movie.SeatAvailable{
					{"A1": true},
					{"A2": true},
					{"A3": true},
					{"B1": true},
					{"B2": true},
					{"B3": true},
					{"C1": true},
					{"C2": true},
					{"C3": true},
					{"D1": true},
					{"D2": true},
					{"D3": true},
				},
			}

			datas = append(datas, data)

		}

		return datas

	}()

	_, err = col.InsertMany(pctx, addMovieAvailable)

	if err != nil {
		log.Printf("Error: Count Documents failed: %v", err)
		return errors.New("error: countDocuments failed")
	}

	log.Println("item id is", result.InsertedID)

	return nil

}

func (r *moviesrepository) FindOneMovie(pctx context.Context, movieId string) (*movie.Movie, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	result := new(movie.Movie)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(movieId)}).Decode(result); err != nil {
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

func (r *moviesrepository) FindComingSoonMovie(pctx context.Context, filter any) ([]*movie.MovieData, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		log.Printf("Error: Find Coming SoonMovie Failed: %s", err.Error())
		return make([]*movie.MovieData, 0), nil
	}

	results := make([]*movie.MovieData, 0)

	for cursor.Next(ctx) {
		result := new(movie.MovieData)
		if err := cursor.Decode(result); err != nil {
			log.Println("Error: Find All Movie Failed:", err.Error())
			return make([]*movie.MovieData, 0), errors.New("error: find all item failed")
		}

		results = append(results, &movie.MovieData{
			MovieId:  result.MovieId,
			Title:    result.Title,
			Price:    result.Price,
			ImageUrl: result.ImageUrl,
		})

	}

	return results, nil
}

func (r *moviesrepository) FindMovieShowtime(pctx context.Context, movieId string) ([]*movie.MovieShowTimeRes, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie_available")

	results := make([]*movie.MovieShowTimeRes, 0)

	log.Println(utils.GetLocaltime())

	cursor, err := col.Find(ctx, bson.M{"movie_id": movieId})
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
			Movie_id:      result.Id.Hex(),
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

	fmt.Println("req.MovieId is ", req.MovieId)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(req.MovieId)}).Decode(result); err != nil {
		log.Printf("Error: Find Seat Status Failed:%s", err.Error())
		return errors.New("error: find seat status failed")
	}

	for _, reserveSeatNo := range req.SeatNo {
		for x, seatAvailable := range result.SeatAvailable {
			fmt.Println("index seat is", x)
			if _, ok := seatAvailable[reserveSeatNo]; ok {
				fmt.Println("Gore Buy Dai Na")
				result.SeatAvailable[x][reserveSeatNo] = false
				break
			} else if x == (len(result.SeatAvailable) - 1) {
				fmt.Println("THis Csae")
				log.Println("error:no seat match")
				return errors.New("error: no seat match")
			}
		}
	}

	updateResult, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(req.MovieId)}, bson.M{"$set": result})
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

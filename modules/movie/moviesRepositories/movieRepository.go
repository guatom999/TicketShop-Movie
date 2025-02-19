package moviesRepositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/guatom999/TicketShop-Movie/pkg/queue"
	"github.com/guatom999/TicketShop-Movie/utils"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/segmentio/kafka-go"
)

type (
	MoviesRepositoryService interface {
		InsertMovie(pctx context.Context, req []*movie.Movie, movieRound int64) error
		FindOneMovie(pctx context.Context, movieId string) (*movie.MovieShowCase, error)
		FindAllMovie(pctx context.Context, filter any) ([]*movie.MovieData, error)
		FindComingSoonMovie(pctx context.Context, filter any) ([]*movie.MovieData, error)
		FindMovieShowtime(pctx context.Context, movieId string) ([]*movie.MovieShowTimeRes, error)
		GetOneMovieAvaliable(pctx context.Context, req *movie.ReserveDetailReq) (*movie.MovieAvaliable, error)
		UpdateSeatStatus(pctx context.Context, movidId string, req *movie.MovieAvaliable) error
		ReserveSeatRes(pctx context.Context, cfg *config.Config, req *movie.ReserveSeatRes) error
	}

	moviesrepository struct {
		db    *mongo.Client
		redis *redis.Client
	}
)

func NewMoviesrepository(db *mongo.Client, redis *redis.Client) MoviesRepositoryService {
	return &moviesrepository{db: db, redis: redis}
}

func MovieProducer(pctx context.Context, cfg *config.Config, topic string) *kafka.Conn {
	conn := queue.KafkaConn(cfg, topic)

	topicConfigs := kafka.TopicConfig{
		Topic:             string(topic),
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	// topicConfigs := make([]kafka.TopicConfig, 0)

	// if !queue.IsTopicIsAlreadyExits(conn, cfg.Kafka.Topic) {
	// 	topicConfigs = append(topicConfigs, kafka.TopicConfig{
	// 		Topic:             topic,
	// 		NumPartitions:     1,
	// 		ReplicationFactor: 1,
	// 	})
	// }

	if err := conn.CreateTopics(topicConfigs); err != nil {
		log.Printf("Erorr: Create Topic Failed %s", err.Error())
		panic(err.Error())
	}

	return conn

}

func (r *moviesrepository) InsertMovie(pctx context.Context, req []*movie.Movie, movieRound int64) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	toAddMovies := func() []any {
		datas := make([]any, 0)

		for i := 0; i < len(req); i++ {
			data := &movie.Movie{
				Title:           req[i].Title,
				Description:     req[i].Description,
				RunningTime:     req[i].RunningTime,
				Price:           req[i].Price,
				ImageUrl:        req[i].ImageUrl,
				CreatedAt:       req[i].CreatedAt,
				UpdatedAt:       req[i].UpdatedAt,
				Category:        req[i].Category,
				ReleaseAt:       req[i].ReleaseAt,
				OutOfTheatersAt: req[i].OutOfTheatersAt,
			}

			datas = append(datas, data)
		}

		return datas

	}()

	result, err := col.InsertMany(ctx, toAddMovies)
	if err != nil {
		log.Printf("Error: Insert One Movie Failed")
		return errors.New("error: insert one movie failed")
	}

	dayLength := req[len(req)-1].OutOfTheatersAt.YearDay() - req[0].ReleaseAt.YearDay() + 1

	lastInsertId := result.InsertedIDs[0].(primitive.ObjectID)

	col = db.Collection("movie_available")

	addMovieAvailable := func() []any {

		datas := make([]any, 0)

		for i := 0; i < len(req); i++ {
			for x := 0; x < dayLength*int(movieRound); x++ {
				data := movie.MovieAvaliable{
					Movie_Id:  result.InsertedIDs[i].(primitive.ObjectID).Hex(),
					Title:     req[i].Title,
					CreatedAt: utils.GetLocaltime(),
					UpdatedAt: func() time.Time {
						return utils.GetLocaltime()
					}(),
					Showtime: utils.SetSpecificTime(req[i].ReleaseAt.Year(), req[i].ReleaseAt.Month(), req[i].ReleaseAt.Day()+int(math.Floor(float64(x)/float64(movieRound))), 10+(x%int(movieRound)), 30, 0),
					//Need Refactor
					SeatAvailable: []movie.SeatAvailable{
						{"A1": true},
						{"A2": true},
						{"A3": true},
						{"A4": true},
						{"A5": true},
						{"A6": true},
						{"A7": true},
						{"A8": true},
						{"A9": true},
						{"A10": true},
						{"A11": true},
						{"A12": true},
						{"B1": true},
						{"B2": true},
						{"B3": true},
						{"B4": true},
						{"B5": true},
						{"B6": true},
						{"B7": true},
						{"B8": true},
						{"B9": true},
						{"B10": true},
						{"B11": true},
						{"B12": true},
						{"C1": true},
						{"C2": true},
						{"C3": true},
						{"C4": true},
						{"C5": true},
						{"C6": true},
						{"C7": true},
						{"C8": true},
						{"C9": true},
						{"C10": true},
						{"C11": true},
						{"C12": true},
						{"D1": true},
						{"D2": true},
						{"D3": true},
						{"D4": true},
						{"D5": true},
						{"D6": true},
						{"D7": true},
						{"D8": true},
						{"D9": true},
						{"D10": true},
						{"D11": true},
						{"D12": true},
					},
				}

				datas = append(datas, data)
			}
		}

		return datas

	}()

	_, err = col.InsertMany(ctx, addMovieAvailable)

	if err != nil {
		log.Printf("Error: Count Documents failed: %v", err)
		return errors.New("error: countDocuments failed")
	}

	log.Println("first item id is", lastInsertId)

	return nil

}

func (r *moviesrepository) FindOneMovie(pctx context.Context, movieId string) (*movie.MovieShowCase, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	result := new(movie.Movie)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(movieId)}).Decode(result); err != nil {
		log.Printf("Error: FindOne Movie Failed:%s", err.Error())
		return nil, errors.New("error: findone movie failed")
	}

	return &movie.MovieShowCase{
		Title:       result.Title,
		Description: result.Description,
		RunningTime: result.RunningTime,
		Price:       result.Price,
		ImageUrl:    result.ImageUrl,
	}, nil
}

func (r *moviesrepository) FindAllMovie(pctx context.Context, filter any) ([]*movie.MovieData, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*30)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	results := make([]*movie.MovieData, 0)

	value, err := r.redis.Get(ctx, "movies_list").Result()
	if err == nil {
		if err = json.Unmarshal([]byte(value), &results); err == nil {
			return results, nil
		}
	}

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		log.Printf("Error: Find All Movie Failed: %s", err.Error())
		return make([]*movie.MovieData, 0), errors.New("error: find all item failed")
	}

	for cursor.Next(ctx) {
		result := new(movie.Movie)
		if err := cursor.Decode(result); err != nil {
			log.Println("Error: Find All Movie Failed:", err.Error())
			return make([]*movie.MovieData, 0), errors.New("error: find all item failed")
		}

		results = append(results, &movie.MovieData{
			MovieId:    result.Id.Hex(),
			Title:      result.Title,
			Release_At: utils.GetStringTime(result.ReleaseAt),
			Price:      result.Price,
			ImageUrl:   result.ImageUrl,
		})
	}

	data, err := json.Marshal(results)
	if err != nil {
		return make([]*movie.MovieData, 0), errors.New("error: find all item failed")
	}

	_, err = r.redis.Set(ctx, "movies_list", string(data), time.Second*30).Result()
	if err != nil {
		return make([]*movie.MovieData, 0), errors.New("error: set redis failed")
	}

	return results, nil
}

func (r *moviesrepository) FindComingSoonMovie(pctx context.Context, filter any) ([]*movie.MovieData, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie")

	results := make([]*movie.MovieData, 0)

	value, err := r.redis.Get(ctx, "comingsoon_list").Result()
	if err == nil {
		if err = json.Unmarshal([]byte(value), &results); err == nil {
			return results, nil
		}
	}

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		log.Printf("Error: Find Coming SoonMovie Failed: %s", err.Error())
		return make([]*movie.MovieData, 0), nil
	}

	for cursor.Next(ctx) {
		result := new(movie.Movie)
		if err := cursor.Decode(result); err != nil {
			log.Println("Error: Find All Movie Failed:", err.Error())
			return make([]*movie.MovieData, 0), errors.New("error: find all item failed")
		}

		results = append(results, &movie.MovieData{
			MovieId:    result.Id.Hex(),
			Title:      result.Title,
			Release_At: utils.GetStringTime(result.ReleaseAt),
			Price:      result.Price,
			ImageUrl:   result.ImageUrl,
		})

	}

	data, err := json.Marshal(results)
	if err != nil {
		return make([]*movie.MovieData, 0), errors.New("error: find comign soon movies failed")
	}

	_, err = r.redis.Set(ctx, "comingsoon_list", string(data), time.Second*30).Result()
	if err != nil {
		return make([]*movie.MovieData, 0), errors.New("error: set redis failed")
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
			IsComingSoon:  true,
		})

	}

	return results, nil
}

func (r *moviesrepository) GetOneMovieAvaliable(pctx context.Context, req *movie.ReserveDetailReq) (*movie.MovieAvaliable, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie_available")

	result := new(movie.MovieAvaliable)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(req.MovieId)}).Decode(result); err != nil {
		log.Printf("Error: Find Seat Status Failed:%s", err.Error())
		return nil, errors.New("error: find seat status failed")
	}

	return result, nil

	// for _, reserveSeatNo := range req.SeatNo {
	// 	for x, seatAvailable := range result.SeatAvailable {
	// 		if _, ok := seatAvailable[reserveSeatNo]; ok {
	// 			result.SeatAvailable[x][reserveSeatNo] = false
	// 			break
	// 		} else if x == (len(result.SeatAvailable) - 1) {
	// 			log.Println("error:no seat match")
	// 			return errors.New("error: no seat match")
	// 		}
	// 	}
	// }

	// updateResult, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(req.MovieId)}, bson.M{"$set": result})
	// if err != nil {
	// 	log.Printf("Error: Update Seat Status Failed %v", err)
	// 	return errors.New("error: update seat status failed")
	// }

	// log.Printf("update status is :%v", updateResult)

	// return nil

}

func (r *moviesrepository) UpdateSeatStatus(pctx context.Context, movidId string, req *movie.MovieAvaliable) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie_available")

	_, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(movidId)}, bson.M{"$set": req})
	if err != nil {
		log.Printf("Error: Update Seat Status Failed %v", err)
		return errors.New("error: update seat status failed")
	}

	// log.Printf("update status is :%v", updateResult)

	return nil
}

func (r *moviesrepository) ReserveSeatRes(pctx context.Context, cfg *config.Config, req *movie.ReserveSeatRes) error {
	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	conn := MovieProducer(ctx, cfg, "reserve-seat-res")

	message := kafka.Message{
		Key:   []byte("payment"),
		Value: utils.EncodeMessage(req),
	}

	// conn.SetWriteDeadline(time.Now().Add(20 * time.Second))
	_, err := conn.WriteMessages(message)

	if err != nil {
		log.Fatal("failed to write messages:", err)
		return errors.New("error: failed to send message")
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
		return errors.New("error: failed to close broker")
	}

	fmt.Println("Send Message Success")

	return nil
}

// func (r *moviesrepository) ReserveSeatRes(pctx context.Context, cfg *config.Config, req *movie.RollBackReserveSeatRes) error {

// 	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
// 	defer cancel()

// 	conn := MovieProducer(ctx, cfg, "roll-back-res")

// 	message := kafka.Message{
// 		Value: utils.EncodeMessage(req),
// 	}

// 	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
// 	_, err := conn.WriteMessages(message)

// 	if err != nil {
// 		log.Fatal("failed to write messages:", err)
// 		return errors.New("error: failed to send message")
// 	}

// 	if err := conn.Close(); err != nil {
// 		log.Fatal("failed to close writer:", err)
// 		return errors.New("error: failed to close broker")
// 	}

// 	fmt.Println("Send Message Success")

// 	return nil
// }

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

func (r *moviesrepository) InsertNews(pctx context.Context, req *movie.AddNewsReq) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie_news")

	_, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: Insert Movie News Failed:%s", err.Error())
		return errors.New("error: insert movie news failed")
	}

	return nil
}

func (r *moviesrepository) InsertPromotions(pctx context.Context, req *movie.AddPromotionsReq) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("movie_db")
	col := db.Collection("movie_promotions")

	_, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: Insert Movie News Failed:%s", err.Error())
		return errors.New("error: insert movie news failed")
	}

	return nil
}

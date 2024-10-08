package migrate

import (
	"context"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"github.com/guatom999/TicketShop-Movie/modules/movie"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MovieMigrate(pctx context.Context, cfg *config.Config) {

	db := database.DbConn(pctx, cfg).Database("movie_db")
	defer db.Client().Disconnect(pctx)

	col := db.Collection("movie")

	// set index
	_, err := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"title", 1}}},
	})

	if err != nil {
		log.Fatalf("Error: CreateIndex Failed")
		panic(err)
	}

	documents := func() []any {
		mockdatas := []*movie.Movie{
			{
				Title:           "Lalaland",
				Description:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. ",
				RunningTime:     "1 ชั่วโมง 30 นาที",
				Price:           150,
				ImageUrl:        "https://i5.walmartimages.com/seo/La-La-Land-Movie-Poster-Poster-Print-24-x-36_20f02811-01b4-4aea-9bb2-a79942bd2642_1.856c035d66f8fd216f6d933259bc3dfb.jpeg",
				CreatedAt:       utils.GetLocaltime(),
				UpdatedAt:       utils.GetLocaltime(),
				Category:        "RomCom",
				ReleaseAt:       utils.GetLocaltime(),
				OutOfTheatersAt: utils.GetLocaltime().Add(time.Hour * 168),
			},
			{
				Title:           "GI JOE",
				Description:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. ",
				RunningTime:     "1 ชั่วโมง 30 นาที",
				Price:           150,
				ImageUrl:        "https://m.media-amazon.com/images/M/MV5BMTQzMTU1NzQwNl5BMl5BanBnXkFtZTcwNDg4NzMzMw@@._V1_.jpg",
				CreatedAt:       utils.GetLocaltime(),
				UpdatedAt:       utils.GetLocaltime(),
				ReleaseAt:       utils.GetLocaltime(),
				Category:        "RomCom",
				OutOfTheatersAt: utils.GetLocaltime().Add(time.Hour * 168),
			},
			{
				Title:           "HarryPotter1",
				Description:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. ",
				RunningTime:     "1 ชั่วโมง 30 นาที",
				Price:           150,
				ImageUrl:        "https://lh3.googleusercontent.com/proxy/Hh8fT3IQBPFWFkHdvwDBig3USKUOCYXqvzcVWq1Rj_S3tm1k0NzUUlbrjjHyWjCylx6bHsvhOhGdQ_EfsRUDYlR86b1TJZZzjtcDAetDy-rsTdDKL5lIfcGyjeiW1b3OMbhRcmEEniv6EYBNhRJPZZggmy1QGJ_SYX-iHfg-K_knSf5H",
				CreatedAt:       utils.GetLocaltime(),
				UpdatedAt:       utils.GetLocaltime(),
				ReleaseAt:       utils.GetLocaltime(),
				Category:        "RomCom",
				OutOfTheatersAt: utils.GetLocaltime().Add(time.Hour * 168),
			},
			{
				Title:           "HarryPotter3",
				Description:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. ",
				RunningTime:     "1 ชั่วโมง 30 นาที",
				Price:           150,
				ImageUrl:        "https://lh3.googleusercontent.com/proxy/Hh8fT3IQBPFWFkHdvwDBig3USKUOCYXqvzcVWq1Rj_S3tm1k0NzUUlbrjjHyWjCylx6bHsvhOhGdQ_EfsRUDYlR86b1TJZZzjtcDAetDy-rsTdDKL5lIfcGyjeiW1b3OMbhRcmEEniv6EYBNhRJPZZggmy1QGJ_SYX-iHfg-K_knSf5H",
				CreatedAt:       utils.GetLocaltime(),
				UpdatedAt:       utils.GetLocaltime(),
				ReleaseAt:       utils.GetLocaltime(),
				Category:        "RomCom",
				OutOfTheatersAt: utils.GetLocaltime().Add(time.Hour * 168),
			},
			{
				Title:           "HarryPotter4",
				Description:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. ",
				RunningTime:     "1 ชั่วโมง 30 นาที",
				Price:           150,
				ImageUrl:        "https://lh3.googleusercontent.com/proxy/Hh8fT3IQBPFWFkHdvwDBig3USKUOCYXqvzcVWq1Rj_S3tm1k0NzUUlbrjjHyWjCylx6bHsvhOhGdQ_EfsRUDYlR86b1TJZZzjtcDAetDy-rsTdDKL5lIfcGyjeiW1b3OMbhRcmEEniv6EYBNhRJPZZggmy1QGJ_SYX-iHfg-K_knSf5H",
				CreatedAt:       utils.GetLocaltime(),
				UpdatedAt:       utils.GetLocaltime(),
				ReleaseAt:       utils.GetLocaltime(),
				Category:        "RomCom",
				OutOfTheatersAt: utils.GetLocaltime().Add(time.Hour * 168),
			},
			{
				Title:           "HarryPotter5",
				Description:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. ",
				RunningTime:     "1 ชั่วโมง 30 นาที",
				Price:           150,
				ImageUrl:        "https://lh3.googleusercontent.com/proxy/Hh8fT3IQBPFWFkHdvwDBig3USKUOCYXqvzcVWq1Rj_S3tm1k0NzUUlbrjjHyWjCylx6bHsvhOhGdQ_EfsRUDYlR86b1TJZZzjtcDAetDy-rsTdDKL5lIfcGyjeiW1b3OMbhRcmEEniv6EYBNhRJPZZggmy1QGJ_SYX-iHfg-K_knSf5H",
				CreatedAt:       utils.GetLocaltime(),
				UpdatedAt:       utils.GetLocaltime(),
				ReleaseAt:       utils.GetLocaltime(),
				Category:        "RomCom",
				OutOfTheatersAt: utils.GetLocaltime().Add(time.Hour * 168),
			},
		}

		docs := make([]any, 0)
		for _, i := range mockdatas {
			docs = append(docs, i)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents)
	if err != nil {
		panic(err)
	}

	col = db.Collection("movie_available")

	_, err = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"movie_id", 1}}},
		{Keys: bson.D{{"title", 1}}},
	})

	documents = func() []any {

		mockdatas := []*movie.MovieAvaliable{
			{
				Movie_Id:  "test000000001",
				Title:     "Lalaland",
				CreatedAt: utils.GetLocaltime(),
				UpdatedAt: utils.GetLocaltime(),
				Showtime:  utils.SetSpecificTime(2024, 2, 19, 10, 30, 0),
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
			},
			{
				Movie_Id:  "test000000002",
				Title:     "Lalaland",
				CreatedAt: utils.GetLocaltime(),
				UpdatedAt: utils.GetLocaltime(),
				Showtime:  utils.SetSpecificTime(2024, 2, 19, 12, 30, 0),
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
			},
			{
				Movie_Id:  "test000000003",
				Title:     "Lalaland",
				CreatedAt: utils.GetLocaltime(),
				UpdatedAt: utils.GetLocaltime(),
				Showtime:  utils.SetSpecificTime(2024, 2, 19, 15, 30, 0),
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
			},
		}

		docs := make([]any, 0)
		for _, i := range mockdatas {
			docs = append(docs, i)
		}

		return docs
	}()

	_, err = col.InsertMany(pctx, documents)

	if err != nil {
		log.Fatalf("Error: CreateIndex Movie_available")
		panic(err)
	}

	col = db.Collection("categories")

	_, err = col.Indexes().CreateOne(pctx, mongo.IndexModel{
		Keys: bson.D{{"name", 1}},
	})

	if err != nil {
		log.Fatalf("Error: CreateIndex Failed")
		panic(err)
	}

	documents = func() []any {
		mockData := []movie.Category{
			{
				Name: "Comedy",
			},
			{
				Name: "Romantic",
			},
			{
				Name: "Rom-Com",
			},
			{
				Name: "Action",
			},
			{
				Name: "Music",
			},
		}

		docs := make([]any, 0)

		for _, i := range mockData {
			docs = append(docs, i)
		}

		return docs

	}()

	_, err = col.InsertMany(pctx, documents)
	if err != nil {
		log.Fatalf("Error: CreateIndex Failed")
		panic(err)
	}

	log.Println("Migrate movies completed:", results)

}

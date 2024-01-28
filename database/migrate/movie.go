package migrate

import (
	"context"
	"log"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/database"
	"github.com/guatom999/TicketShop-Movie/modules/movie"
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
				Title:     "Lalaland",
				Price:     150,
				ImageUrl:  "https://i5.walmartimages.com/seo/La-La-Land-Movie-Poster-Poster-Print-24-x-36_20f02811-01b4-4aea-9bb2-a79942bd2642_1.856c035d66f8fd216f6d933259bc3dfb.jpeg",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Avaliable: 15,
			},
			{
				Title:     "GI JOE",
				Price:     150,
				ImageUrl:  "https://m.media-amazon.com/images/M/MV5BMTQzMTU1NzQwNl5BMl5BanBnXkFtZTcwNDg4NzMzMw@@._V1_.jpg",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Avaliable: 15,
			},
			{
				Title:     "HarryPotter1",
				Price:     150,
				ImageUrl:  "https://lh3.googleusercontent.com/proxy/Hh8fT3IQBPFWFkHdvwDBig3USKUOCYXqvzcVWq1Rj_S3tm1k0NzUUlbrjjHyWjCylx6bHsvhOhGdQ_EfsRUDYlR86b1TJZZzjtcDAetDy-rsTdDKL5lIfcGyjeiW1b3OMbhRcmEEniv6EYBNhRJPZZggmy1QGJ_SYX-iHfg-K_knSf5H",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Avaliable: 15,
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

	col = db.Collection("categories")

	_, err = col.Indexes().CreateOne(pctx, mongo.IndexModel{
		Keys: bson.D{{"Name", 1}},
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

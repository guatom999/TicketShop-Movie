package bookRepositories

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	BookRepositoryService interface {
	}

	bookRepository struct {
		db *mongo.Client
	}
)

func NewBookRepository() BookRepositoryService {
	return &bookRepository{}
}

func (r *bookRepository) ConnectBookingDb() *mongo.Database {
	return r.db.Database("booking_db")
}

func (r *bookRepository) FindMovieIsAvaliable(title string) error {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf("http://localhost:8090/getmovie/%s", title))
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	fmt.Println("Response status code:", resp.StatusCode)
	fmt.Println("Response body:", string(body)) // Assuming JSON response

	return nil

}

func (r *bookRepository) ReserveSeat(pctx context.Context) {

}

func (r *bookRepository) BuyTicket(pctx context.Context, title string) error {

	return nil

}

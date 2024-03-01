package customerRepositories

import (
	"context"
)

type (
	CustomerRepositoryService interface {
	}

	customerRepository struct {
	}
)

func NewCustomerRepository() CustomerRepositoryService {

	return &customerRepository{}

}

func (r *customerRepository) AddTicketCustomer(pctx context.Context) error {

	// ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	// defer cancel()

	// db := r.db.Database("movie_db")
	// col := db.Collection("movie")

	return nil

}

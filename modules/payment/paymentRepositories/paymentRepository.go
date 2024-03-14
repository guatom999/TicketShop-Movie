package paymentRepositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PaymentRepositoryService interface {
		BuyItem(pctx context.Context) error
	}

	paymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepositoryService {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) BuyItem(pctx context.Context) error {

	return nil
}

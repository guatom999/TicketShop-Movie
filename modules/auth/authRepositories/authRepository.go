package authRepositories

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/modules/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepositoryService interface {
		Register(pctx context.Context, req *auth.RegisterRequest) (primitive.ObjectID, error)
	}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository() AuthRepositoryService {
	return &authRepository{}
}

func (r *authRepository) Register(pctx context.Context, req *auth.RegisterRequest) (primitive.ObjectID, error) {

	// ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	// defer cancel()

	// db := r.db.Database("")

	return primitive.NilObjectID, nil
}

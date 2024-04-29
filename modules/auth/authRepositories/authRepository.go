package authRepositories

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/modules/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	AuthRepositoryService interface {
		Register(pctx context.Context, req *auth.RegisterRequest) (primitive.ObjectID, error)
	}

	authRepository struct {
	}
)

func NewAuthRepository() AuthRepositoryService {
	return &authRepository{}
}

func (r *authRepository) Register(pctx context.Context, req *auth.RegisterRequest) (primitive.ObjectID, error) {
	return primitive.NilObjectID, nil
}

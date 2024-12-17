package middlewareRepositories

import (
	"context"
	"time"

	"github.com/guatom999/TicketShop-Movie/modules/customer"
)

type (
	IMiddlewareRepositoryService interface {
	}

	middlwareRepository struct {
	}
)

func NewMiddlewareRepository() IMiddlewareRepositoryService {
	return &middlwareRepository{}
}

func (r *middlwareRepository) AccessTokenSearch(pctx context.Context) (*customer.Credential, error) {

	_, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	return nil, nil
}

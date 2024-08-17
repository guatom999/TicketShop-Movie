package middlewareRepositories

import (
	"context"
	"time"
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

func (r *middlwareRepository) AccessTokenSearch(pctx context.Context) error {

	_, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	return nil
}

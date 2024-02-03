package bookRepositories

import "go.mongodb.org/mongo-driver/mongo"

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

// func BookMovie(pctx context.Context, req )

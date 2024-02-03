package bookUseCases

import "github.com/guatom999/TicketShop-Movie/modules/book/bookRepositories"

type (
	BookUseCaseService interface {
	}

	bookUsecase struct {
		bookRepo bookRepositories.BookRepositoryService
	}
)

func NewBookUsecase(bookRepo bookRepositories.BookRepositoryService) BookUseCaseService {
	return &bookUsecase{bookRepo: bookRepo}
}

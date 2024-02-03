package bookHandlers

import (
	"github.com/guatom999/TicketShop-Movie/modules/book/bookUseCases"
)

type (
	BookHandlerService interface {
	}

	bookHandler struct {
		bookUseCases bookUseCases.BookUseCaseService
	}
)

func NewBookUsecase(bookUseCases bookUseCases.BookUseCaseService) BookHandlerService {
	return &bookHandler{bookUseCases: bookUseCases}
}

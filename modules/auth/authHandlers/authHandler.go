package authHandlers

import "github.com/guatom999/TicketShop-Movie/modules/auth/authUseCases"

type (
	AuthHandlerService interface {
	}

	authHandler struct {
		authUseCase authUseCases.AuthUseCaseService
	}
)

func NewAuthHandler(authUseCase authUseCases.AuthUseCaseService) AuthHandlerService {
	return &authHandler{authUseCase}
}

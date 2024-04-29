package authUseCases

import "github.com/guatom999/TicketShop-Movie/modules/auth/authRepositories"

type (
	AuthUseCaseService interface {
	}

	authUseCase struct {
		authRepo authRepositories.AuthRepositoryService
	}
)

func NewUuthUseCase(authRepo authRepositories.AuthRepositoryService) AuthUseCaseService {
	return &authUseCase{authRepo: authRepo}
}

package customerUseCases

type (
	CustomerUseCaseService interface {
	}

	customerUseCase struct {
	}
)

func NewCustomerUseCase() CustomerUseCaseService {
	return &customerUseCase{}
}

package middlewareUseCases

type (
	MiddlewareUserCaseInterface interface {
	}

	middlewareUseCase struct {
	}
)

func NewMiddlwareUseCase() MiddlewareUserCaseInterface {
	return &middlewareUseCase{}
}

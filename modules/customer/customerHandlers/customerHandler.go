package customerHandlers

type (
	CustomerHandlerService interface {
	}

	customerHandler struct {
	}
)

func NewCustomerHandler() CustomerHandlerService {
	return &customerHandler{}
}

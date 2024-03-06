package ticketUseCases

type (
	TicketUseCaseService interface {
	}

	ticketUseCase struct {
	}
)

func NewTicketUseCase() TicketUseCaseService {
	return &ticketUseCase{}
}

package ticketHandlers

type (
	TicketHandlerService interface {
	}

	ticketHandler struct {
	}
)

func NewTicketHandler() TicketHandlerService {
	return &ticketHandler{}
}

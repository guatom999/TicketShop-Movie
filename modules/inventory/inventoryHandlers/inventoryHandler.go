package inventoryHandlers

type (
	InventoryHandlerService interface {
	}

	inventoryHandler struct {
	}
)

func NewInventoryHandler() InventoryHandlerService {
	return &inventoryHandler{}
}

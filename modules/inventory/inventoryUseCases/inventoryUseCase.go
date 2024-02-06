package inventoryUseCases

type (
	InventoryUseCaseService interface {
	}

	inventoryUseCase struct {
	}
)

func NewInventoryUseCase() InventoryUseCaseService {
	return &inventoryUseCase{}
}

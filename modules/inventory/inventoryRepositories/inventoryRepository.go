package inventoryRepositories

type (
	InventoryRepositoryService interface {
	}

	inventoryRepository struct {
	}
)

func NewInventoryRepository() InventoryRepositoryService {
	return &inventoryRepository{}
}

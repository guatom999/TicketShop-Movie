package bookRepositories

type (
	BookRepositoryService interface {
	}

	bookRepository struct {
	}
)

func NewBookRepository() BookRepositoryService {
	return &bookRepository{}
}

package customerRepositories

type (
	CustomerRepositoryService interface {
	}

	customerRepository struct {
	}
)

func NewCustomerRepository() CustomerRepositoryService {

	return &customerRepository{}

}

// func AddPlayerTicket(pctx context.Context, req )

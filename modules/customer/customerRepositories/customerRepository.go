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

// func (r *customerRepository) AddTicketCustomer(pctx context.Context, req *customer.)

package server

import (
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerUseCases"
)

func (s *server) CustomerModules() {
	customerRepo := customerRepositories.NewCustomerRepository(s.db)
	customerUseCase := customerUseCases.NewCustomerUseCase(customerRepo)
	customerHandler := customerHandlers.NewCustomerHandler(customerUseCase)

	userRouter := s.app.Group("/user")

	userRouter.POST("/register", customerHandler.Register)

}

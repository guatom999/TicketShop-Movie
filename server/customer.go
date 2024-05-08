package server

import (
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerUseCases"
)

func (s *server) CustomerModules() {
	customerRepo := customerRepositories.NewCustomerRepository(s.db)
	customerUseCase := customerUseCases.NewCustomerUseCase(customerRepo, s.cfg)
	customerHandler := customerHandlers.NewCustomerHandler(customerUseCase)

	customerRouter := s.app.Group("/user")

	customerRouter.POST("/login", customerHandler.Login)
	customerRouter.POST("/register", customerHandler.Register)

}

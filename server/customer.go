package server

import (
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerUseCases"
)

func (s *server) CustomerModules() {
	customerRepo := customerRepositories.NewCustomerRepository(s.db, s.cfg)
	customerUseCase := customerUseCases.NewCustomerUseCase(customerRepo, s.cfg)
	customerHandler := customerHandlers.NewCustomerHandler(customerUseCase, s.cfg)

	customerRouter := s.app.Group("/user")

	// customerRouter.GET("/test-token" , )

	customerRouter.GET("/testjwt", customerHandler.TestJwtAuthorize, customerHandler.TestMilddeware)
	// customerRouter.GET("/testjwt", customerHandler.TestJwtAuthorize)
	customerRouter.POST("/login", customerHandler.Login)
	customerRouter.POST("/logout", customerHandler.Logout)
	customerRouter.POST("/refresh-token", customerHandler.RefreshToken)
	customerRouter.POST("/register", customerHandler.Register)

}

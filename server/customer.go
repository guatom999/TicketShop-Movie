package server

import (
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerUseCases"
	"github.com/guatom999/TicketShop-Movie/modules/middlewares/middlewareHandlers"
)

func (s *server) CustomerModules(
	authMiddleware middlewareHandlers.MiddlewareHandlerInterface,
) {
	customerRepo := customerRepositories.NewCustomerRepository(s.db, s.cfg)
	customerUseCase := customerUseCases.NewCustomerUseCase(customerRepo, s.cfg, s.mailer)
	customerHandler := customerHandlers.NewCustomerHandler(customerUseCase, s.cfg)

	customerRouter := s.app.Group("/user")

	// customerRouter.GET("/test-token" , )

	customerRouter.GET("/testjwt", customerHandler.TestJwtAuthorize, authMiddleware.JwtAuthorize)
	// customerRouter.GET("/testjwt", customerHandler.TestJwtAuthorize, m.JwtAuthorize)
	// customerRouter.GET("/testjwt", customerHandler.TestJwtAuthorize)
	customerRouter.POST("/find-access-token", customerHandler.FindAccessToken)
	customerRouter.POST("/login", customerHandler.Login)
	customerRouter.POST("/logout", customerHandler.Logout)
	customerRouter.POST("/refresh-token", customerHandler.RefreshToken)
	customerRouter.POST("/register", customerHandler.Register)

	customerRouter.POST("/testsendemail", customerHandler.TestSendEmail)

}

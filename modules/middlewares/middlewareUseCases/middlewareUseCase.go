package middlewareUseCases

import (
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/middlewares/middlewareRepositories"
	"github.com/guatom999/TicketShop-Movie/pkg/jwtauth"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareUserCaseInterface interface {
		JwtAuthorize(c echo.Context, accessToken string) (echo.Context, error)
	}

	middlewareUseCase struct {
		middlewareRepository middlewareRepositories.IMiddlewareRepositoryService
		cfg                  *config.Config
	}
)

func NewMiddlwareUseCase(
	middlewareRepository middlewareRepositories.IMiddlewareRepositoryService,
	cfg *config.Config,
) MiddlewareUserCaseInterface {
	return &middlewareUseCase{
		middlewareRepository: middlewareRepository,
		cfg:                  cfg,
	}
}

func (u *middlewareUseCase) JwtAuthorize(c echo.Context, accessToken string) (echo.Context, error) {

	ctx := c.Request().Context()

	claims, err := jwtauth.ParseToken(u.cfg.Jwt.AccessSecretKey, accessToken)
	if err != nil {
		log.Printf("Erorr is %s", err.Error())
		return nil, err
	}

	if err := u.middlewareRepository.AccessTokenSearch(ctx, accessToken); err != nil {
		return nil, err
	}

	c.Set("customer_id", claims.CustomerId)

	return c, nil
}

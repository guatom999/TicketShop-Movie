package middlewareHandlers

import (
	"net/http"
	"strings"

	"github.com/guatom999/TicketShop-Movie/modules/middlewares/middlewareUseCases"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareHandlerInterface interface {
		JwtAuthorize(next echo.HandlerFunc) echo.HandlerFunc
	}

	middlewareHandler struct {
		middlewareUsecase middlewareUseCases.MiddlewareUserCaseInterface
	}
)

func NewMiddlewareHandler(
	middlewareUsecase middlewareUseCases.MiddlewareUserCaseInterface,
) MiddlewareHandlerInterface {
	return &middlewareHandler{middlewareUsecase: middlewareUsecase}
}

func (h *middlewareHandler) JwtAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

		newCtx, err := h.middlewareUsecase.JwtAuthorize(c, accessToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		return next(newCtx)

	}
}

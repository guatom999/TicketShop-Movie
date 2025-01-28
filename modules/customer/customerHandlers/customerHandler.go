package customerHandlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/customer"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerUseCases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type (
	CustomerHandlerService interface {
		Login(c echo.Context) error
		RefreshToken(c echo.Context) error
		Logout(c echo.Context) error
		GetCustomerProfile(c echo.Context) error
		TestMilddeware(next echo.HandlerFunc) echo.HandlerFunc
		TestJwtAuthorize(c echo.Context) error
		Register(c echo.Context) error
		FindAccessToken(c echo.Context) error
		ForgotPassword(c echo.Context) error
	}

	customerHandler struct {
		customerUseCase customerUseCases.CustomerUseCaseService
		cfg             *config.Config
	}
)

func NewCustomerHandler(customerUseCase customerUseCases.CustomerUseCaseService, cfg *config.Config) CustomerHandlerService {
	return &customerHandler{
		customerUseCase: customerUseCase,
		cfg:             cfg,
	}
}

func (h *customerHandler) GetCustomerProfile(c echo.Context) error {

	ctx := context.Background()

	customerId := c.Param("customer_id")

	result, err := h.customerUseCase.GetCustomerProfile(ctx, customerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}
func (h *customerHandler) Login(c echo.Context) error {

	ctx := context.Background()

	req := new(customer.LoginReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "something weng wrong")
	}

	res, err := h.customerUseCase.Login(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *customerHandler) TestJwtAuthorize(c echo.Context) error {

	return c.JSON(http.StatusOK, "test success")
}

func (h *customerHandler) RefreshToken(c echo.Context) error {

	ctx := context.Background()

	req := new(customer.CustomerRefreshTokenReq)

	if err := c.Bind(req); err != nil {
		log.Printf("req wrong cause of :++++>", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.customerUseCase.RefreshToken(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *customerHandler) Logout(c echo.Context) error {

	ctx := context.Background()

	req := new(customer.LogoutReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.customerUseCase.Logout(ctx, req.CredentialId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Logout success Deleted user count %d", res))
}

func (h *customerHandler) FindAccessToken(c echo.Context) error {

	ctx := context.Background()

	req := new(customer.FindAccessTokenReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.customerUseCase.FindAccessToken(ctx, req.AccessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (h *customerHandler) TestMilddeware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

		result, err := h.customerUseCase.TestMiddleware(c, accessToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		return next(result)

	}
}

func (h *customerHandler) Register(c echo.Context) error {

	ctx := context.Background()

	req := new(customer.RegisterReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.customerUseCase.Register(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *customerHandler) ForgotPassword(c echo.Context) error {

	ctx := context.Background()

	req := new(customer.SendForgotPasswordReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.customerUseCase.ForgotPassword(ctx, req.Email); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "send email success")
}

package customerUseCases

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/customer"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerRepositories"
	"github.com/guatom999/TicketShop-Movie/pkg/jwtauth"
	"github.com/guatom999/TicketShop-Movie/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type (
	CustomerUseCaseService interface {
		Login(pctx context.Context, req *customer.LoginReq) (*customer.CustomerProfileRes, error)
		Logout(pctx context.Context, credentialId string) (int64, error)
		GetCustomerProfile(pctx context.Context, customerId string) (*customer.CustomerProfile, error)
		FindAccessToken(pctx context.Context, accessToken string) (*customer.Credential, error)
		Register(pctx context.Context, req *customer.RegisterReq) (primitive.ObjectID, error)
		RefreshToken(pctx context.Context, req *customer.CustomerRefreshTokenReq) (*customer.CustomerProfileRes, error)
		TestMiddleware(c echo.Context, accessToken string) (echo.Context, error)
		TestSendEmail(pctx context.Context, sendTo string) error
	}

	customerUseCase struct {
		customerRepo customerRepositories.CustomerRepositoryService
		cfg          *config.Config
		mailer       *gomail.Dialer
	}
)

func NewCustomerUseCase(customerRepo customerRepositories.CustomerRepositoryService, cfg *config.Config, mailer *gomail.Dialer) CustomerUseCaseService {
	return &customerUseCase{
		customerRepo: customerRepo,
		cfg:          cfg,
		mailer:       mailer,
	}
}

func (u *customerUseCase) Login(pctx context.Context, req *customer.LoginReq) (*customer.CustomerProfileRes, error) {

	result, err := u.customerRepo.FindOneCustomerWithCredential(pctx, req.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	customerId := "customer:" + result.Id.Hex()

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("error: password mismatch")
	}

	accessToken := u.customerRepo.NewAccessToken(u.cfg, &customer.Claims{
		Id: result.Id.Hex(),
	})

	refreshToken := u.customerRepo.NewRefreshToken(u.cfg, &customer.Claims{
		Id: result.Id.Hex(),
	})

	credential, _ := u.customerRepo.InsertCustomerCredential(pctx, &customer.Credential{
		CustomerId:   customerId,
		Rolecode:     1,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	customerCredential, err := u.customerRepo.FindCustomerCredential(pctx, credential.Hex())
	if err != nil {
		return nil, err
	}

	// loc, _ := time.LoadLocation("Asia/Bangkok")

	return &customer.CustomerProfileRes{
		Status: "ok",
		CustomerProfile: &customer.CustomerProfile{
			Id:         credential.Hex(),
			CustomerId: customerId,
			Email:      result.Email,
			ImageUrl:   result.ImageUrl,
			UserName:   result.UserName,
			Created_At: utils.GetStringTime(result.Created_At),
			Updated_At: utils.GetStringTime(result.Updated_At),
			Credential: &customer.CredentailRes{
				AccessToken:  customerCredential.AccessToken,
				RefreshToken: customerCredential.RefreshToken,
			},
		},
	}, nil
}

func (u *customerUseCase) Logout(pctx context.Context, credentialId string) (int64, error) {

	return u.customerRepo.DeleteCustomerCredential(pctx, credentialId)
}

// func (u *customerUseCase) GetCustomerProfile(pctx context.Context, customerId string) (*customer.CustomerProfile, error) {

// 	result, err := u.customerRepo.FindCustomer(pctx, customerId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &customer.CustomerProfile{
// 		Id:         result.Id.Hex(),
// 		CustomerId: customerId,
// 		Email:      result.Email,
// 		ImageUrl:   result.ImageUrl,
// 		UserName:   result.UserName,
// 		Created_At: utils.GetStringTime(result.Created_At),
// 		Updated_At: utils.GetStringTime(result.Updated_At),
// 	}, nil
// }

func (u *customerUseCase) GetCustomerProfile(pctx context.Context, customerId string) (*customer.CustomerProfile, error) {

	result, err := u.customerRepo.FindCustomer(pctx, customerId)
	if err != nil {
		return nil, err
	}

	return &customer.CustomerProfile{
		Id:         result.Id.Hex(),
		CustomerId: customerId,
		Email:      result.Email,
		ImageUrl:   result.ImageUrl,
		UserName:   result.UserName,
		Created_At: utils.GetStringTime(result.Created_At),
		Updated_At: utils.GetStringTime(result.Updated_At),
	}, nil
}

func (u *customerUseCase) RefreshToken(pctx context.Context, req *customer.CustomerRefreshTokenReq) (*customer.CustomerProfileRes, error) {

	_, err := jwtauth.ParseToken(u.cfg.Jwt.RefreshSecretKey, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	fmt.Println("after trim Prefix is", strings.TrimPrefix(req.CustomerId, "customer:"))

	customerProfile, err := u.customerRepo.FindCustomer(pctx, strings.TrimPrefix(req.CustomerId, "customer:"))
	if err != nil {
		return nil, err
	}

	customerId := "customer:" + customerProfile.Id.Hex()

	accessToken := u.customerRepo.NewAccessToken(u.cfg, &customer.Claims{
		Id: customerId,
	})

	refreshToken := u.customerRepo.ReloadToken(u.cfg, &customer.Claims{
		Id: customerId,
	})

	if err := u.customerRepo.UpdateCustomerCredential(pctx, req.CredentialId, &customer.UpdateRefreshToken{
		CustomerId:   customerProfile.Id.Hex(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UpdatedAt:    utils.GetLocaltime(),
	}); err != nil {
		return nil, err
	}

	customerCredential, err := u.customerRepo.FindCustomerCredential(pctx, req.CredentialId)
	if err != nil {
		return nil, err
	}

	return &customer.CustomerProfileRes{
		Status: "ok",
		CustomerProfile: &customer.CustomerProfile{
			Id:         customerProfile.Id.Hex(),
			CustomerId: customerId,
			Email:      customerProfile.Email,
			ImageUrl:   customerProfile.ImageUrl,
			UserName:   customerProfile.UserName,
			Created_At: utils.GetStringTime(customerProfile.Created_At),
			Updated_At: utils.GetStringTime(customerProfile.Updated_At),
			Credential: &customer.CredentailRes{
				AccessToken:  customerCredential.AccessToken,
				RefreshToken: customerCredential.RefreshToken,
			},
		},
	}, nil
}

func (u *customerUseCase) FindAccessToken(pctx context.Context, accessToken string) (*customer.Credential, error) {
	return u.customerRepo.FindAccessToken(pctx, accessToken)
}

func (u *customerUseCase) TestMiddleware(c echo.Context, accessToken string) (echo.Context, error) {

	ctx := c.Request().Context()

	claims, err := jwtauth.ParseToken(u.cfg.Jwt.AccessSecretKey, accessToken)
	if err != nil {
		log.Printf("Erorr is %s", err.Error())
		return nil, err
	}

	result, err := u.customerRepo.FindAccessToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("errro: access token is invalid")
	}

	// fmt.Println("customer_id is", claims.Id)
	c.Set("customer_id", claims.CustomerId)

	return c, nil
}
func (u *customerUseCase) Register(pctx context.Context, req *customer.RegisterReq) (primitive.ObjectID, error) {

	if u.customerRepo.IsUserAlreadyExist(pctx, req.UserName, req.Email) {
		return primitive.NilObjectID, errors.New("error: user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Printf("Error: hashedPassword failed %s", err.Error())
		return primitive.NilObjectID, errors.New("error: something wrong with password")
	}

	result, err := u.customerRepo.InsertCustomer(pctx, &customer.Customer{
		UserName:   req.UserName,
		Email:      req.Email,
		ImageUrl:   "https://t4.ftcdn.net/jpg/05/49/98/39/360_F_549983970_bRCkYfk0P6PP5fKbMhZMIb07mCJ6esXL.jpg",
		Password:   string(hashedPassword),
		Created_At: utils.GetLocaltime(),
		Updated_At: utils.GetLocaltime(),
	})
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result, nil
}

func (u *customerUseCase) TestSendEmail(pctx context.Context, sendTo string) error {

	// if err := utils.SendEmail(u.cfg, "bebeoblybe@gmail.com", "test", "justTest"); err != nil {
	// 	log.Printf("Errors TestSendEmail is %s", err.Error())
	// 	return err
	// }

	toSendMessage := gomail.NewMessage()
	toSendMessage.SetHeader("To", sendTo)
	toSendMessage.SetHeader("Subject", "Why you forgot password")
	toSendMessage.SetBody("text/html", "You password is")

	if err := utils.SecondSendEmail(u.cfg, toSendMessage); err != nil {
		log.Printf("Errors TestSendEmail is %s", err.Error())
		return err
	}

	return nil
}

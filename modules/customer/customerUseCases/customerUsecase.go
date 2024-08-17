package customerUseCases

import (
	"context"
	"errors"
	"log"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/customer"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerRepositories"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type (
	CustomerUseCaseService interface {
		Login(pctx context.Context, req *customer.LoginReq) (*customer.CustomerProfileRes, error)
		Register(pctx context.Context, req *customer.RegisterReq) (primitive.ObjectID, error)
	}

	customerUseCase struct {
		customerRepo customerRepositories.CustomerRepositoryService
		cfg          *config.Config
	}
)

func NewCustomerUseCase(customerRepo customerRepositories.CustomerRepositoryService, cfg *config.Config) CustomerUseCaseService {
	return &customerUseCase{
		customerRepo: customerRepo,
		cfg:          cfg,
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

	accessToken := u.customerRepo.AccessToken(u.cfg, &customer.Claims{
		Id:       result.Id.Hex(),
		UserName: result.UserName,
	})

	refreshToken := u.customerRepo.RefreshToken(u.cfg, &customer.Claims{
		Id:       result.Id.Hex(),
		UserName: result.UserName,
	})

	u.customerRepo.InsertCustomerCredential(pctx, &customer.Credential{
		CustomerId:   customerId,
		Rolecode:     1,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	return &customer.CustomerProfileRes{
		Status: "ok",
		CustomerProfile: &customer.CustomerProfile{
			Id:         result.Id.Hex(),
			Email:      result.Email,
			ImageUrl:   result.ImageUrl,
			UserName:   result.UserName,
			Created_At: utils.GetStringTime(result.Created_At),
			Updated_At: utils.GetStringTime(result.Updated_At),
			Credential: &customer.CredentailRes{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		},
	}, nil
}

func (u *customerUseCase) Register(pctx context.Context, req *customer.RegisterReq) (primitive.ObjectID, error) {

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
		return result, err
	}

	return result, nil
}

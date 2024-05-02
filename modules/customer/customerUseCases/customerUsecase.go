package customerUseCases

import (
	"context"
	"errors"
	"log"

	"github.com/guatom999/TicketShop-Movie/modules/customer"
	"github.com/guatom999/TicketShop-Movie/modules/customer/customerRepositories"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type (
	CustomerUseCaseService interface {
		Register(pctx context.Context, req *customer.RegisterReq) (primitive.ObjectID, error)
	}

	customerUseCase struct {
		customerRepo customerRepositories.CustomerRepositoryService
	}
)

func NewCustomerUseCase(customerRepo customerRepositories.CustomerRepositoryService) CustomerUseCaseService {
	return &customerUseCase{customerRepo: customerRepo}
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
		Password:   string(hashedPassword),
		Created_At: utils.GetLocaltime(),
		Updated_At: utils.GetLocaltime(),
	})
	if err != nil {
		return result, err
	}

	return result, nil
}

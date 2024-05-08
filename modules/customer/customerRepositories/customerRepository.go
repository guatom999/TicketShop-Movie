package customerRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/customer"
	"github.com/guatom999/TicketShop-Movie/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	CustomerRepositoryService interface {
		FindOneCustomerWithCredential(pctx context.Context, email string) (*customer.Customer, error)
		InsertCustomer(pctx context.Context, req *customer.Customer) (primitive.ObjectID, error)
		AccessToken(cfg *config.Config, customerPassport *customer.Claims) string
		RefreshToken(cfg *config.Config, customerPassport *customer.Claims) string
	}

	customerRepository struct {
		db  *mongo.Client
		cfg *config.Config
	}
)

func NewCustomerRepository(db *mongo.Client) CustomerRepositoryService {

	return &customerRepository{db: db}

}

func (r *customerRepository) AddTicketCustomer(pctx context.Context) error {

	// ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	// defer cancel()

	// db := r.db.Database("movie_db")
	// col := db.Collection("movie")

	return nil

}

func (r *customerRepository) FindOneCustomerWithCredential(pctx context.Context, email string) (*customer.Customer, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer")

	result := new(customer.Customer)

	if err := col.FindOne(ctx, bson.M{"email": email}).Decode(result); err != nil {
		log.Printf("Error: FindOneCustomerWithCredential Failed %s", err.Error())
		return nil, errors.New("error: find customer failed")
	}

	return result, nil
}

func (r *customerRepository) AccessToken(cfg *config.Config, customerPassport *customer.Claims) string {

	claims := customer.AuthClaims{
		Claims: &customer.Claims{
			Id:       customerPassport.Id,
			UserName: customerPassport.UserName,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "seeyarnmirttown.com",
			Subject:   "access-token",
			Audience:  []string{"seeyarnmirttown.com"},
			ExpiresAt: jwt.NewNumericDate(utils.GetLocaltime().Add(time.Second * 20)),
			NotBefore: jwt.NewNumericDate(utils.GetLocaltime()),
			IssuedAt:  jwt.NewNumericDate(utils.GetLocaltime()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(cfg.Jwt.AccessSecretKey))
	if err != nil {
		log.Printf("Error: SignedToken Failed %s", err.Error())
		return err.Error()
	}

	return accessToken
}

func (r *customerRepository) RefreshToken(cfg *config.Config, customerPassport *customer.Claims) string {
	claims := customer.AuthClaims{
		Claims: &customer.Claims{
			Id:       customerPassport.Id,
			UserName: customerPassport.UserName,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "seeyarnmirttown.com",
			Subject:   "refresh-token",
			Audience:  []string{"seeyarnmirttown.com"},
			ExpiresAt: jwt.NewNumericDate(utils.GetLocaltime().Add(time.Second * 60)),
			NotBefore: jwt.NewNumericDate(utils.GetLocaltime()),
			IssuedAt:  jwt.NewNumericDate(utils.GetLocaltime()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString([]byte(cfg.Jwt.RefreshSecretKey))
	if err != nil {
		log.Printf("Error: SignedToken Failed %s", err.Error())
		return err.Error()
	}

	return refreshToken
}

func (r *customerRepository) InsertCustomer(pctx context.Context, req *customer.Customer) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer")

	customerId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: error insert one customer failed %s", err)
		return primitive.NilObjectID, errors.New("error: register failed")
	}

	return customerId.InsertedID.(primitive.ObjectID), nil
}

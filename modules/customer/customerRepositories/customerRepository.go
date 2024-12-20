package customerRepositories

import (
	"context"
	"errors"
	"fmt"
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
		FindCustomerRefreshToken(pctx context.Context, customerId string) (*customer.Customer, error)
		InsertCustomerCredential(pctx context.Context, req *customer.Credential) (primitive.ObjectID, error)
		DeleteCustomerCredential(pctx context.Context, credentialId string) (int64, error)
		FindAccessToken(pctx context.Context, accessToken string) (*customer.Credential, error)
		FindCustomerCredential(pctx context.Context, credentialId string) (*customer.Credential, error)
		UpdateCustomerCredential(pctx context.Context, credentialId string, req *customer.UpdateRefreshToken) error
		NewAccessToken(cfg *config.Config, customerPassport *customer.Claims) string
		NewRefreshToken(cfg *config.Config, customerPassport *customer.Claims) string
		ReloadToken(cfg *config.Config, customerPassport *customer.Claims) string
	}

	customerRepository struct {
		db  *mongo.Client
		cfg *config.Config
	}
)

func NewCustomerRepository(db *mongo.Client, cfg *config.Config) CustomerRepositoryService {

	return &customerRepository{db: db, cfg: cfg}

}

func (r *customerRepository) AddTicketCustomer(pctx context.Context) error {

	// ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	// defer cancel()

	// db := r.db.Database("movie_db")
	// col := db.Collection("movie")

	return nil

}

func (r *customerRepository) FindOneCustomerCredential(pctx context.Context, customerId string) (*customer.Credential, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer_auth")

	result := new(customer.Credential)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(customerId)}).Decode(result); err != nil {
		log.Printf("Error: FindOneCustomerCredential Failed %s", err.Error())
		return nil, errors.New("error: find credentail customer failed")
	}

	return result, nil
}

func (r *customerRepository) InsertCustomerCredential(pctx context.Context, req *customer.Credential) (primitive.ObjectID, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	req.CreatedAt = utils.GetLocaltime()
	req.UpdatedAt = utils.GetLocaltime()

	db := r.db.Database("customer_db")
	col := db.Collection("customer_auth")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertCustomerCredential Failed %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert credentail customer failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *customerRepository) DeleteCustomerCredential(pctx context.Context, credentialId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer_auth")

	result, err := col.DeleteOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(credentialId)})
	if err != nil {
		log.Printf("Error: DeleteCustomerCredential Failed %s", err.Error())
		return 0, errors.New("error: delete customer credential failed")
	}

	return result.DeletedCount, nil
}

func (r *customerRepository) FindCustomerCredential(pctx context.Context, credentialId string) (*customer.Credential, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer_auth")

	result := new(customer.Credential)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(credentialId)}).Decode(result); err != nil {
		log.Printf("Error: FindCustomerCredential Failed %s", err.Error())
		return nil, errors.New("error: find  customer credential failed")
	}

	return result, nil
}

func (r *customerRepository) UpdateCustomerCredential(pctx context.Context, credentialId string, req *customer.UpdateRefreshToken) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer_auth")

	_, err := col.UpdateOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(credentialId)}, bson.M{"$set": bson.M{
		// "customer_id":   "customer:" + req.CustomerId,
		"access_token":  req.AccessToken,
		"refresh_token": req.RefreshToken,
		"updated_at":    req.UpdatedAt,
	}})

	if err != nil {
		log.Printf("Error: UpdateCustomerCredential Failed %s", err.Error())
		return errors.New("error: update customer credential failed")
	}

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

	fmt.Println("result is", result)

	return result, nil
}

func (r *customerRepository) FindCustomerRefreshToken(pctx context.Context, customerId string) (*customer.Customer, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer")

	result := new(customer.Customer)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertStringToObjectId(customerId)}).Decode(result); err != nil {
		log.Printf("Error: FindCustomerRefreshToken Failed %s", err.Error())
		return nil, errors.New("error: find customer for refresh failed")
	}

	return result, nil
}

func (r *customerRepository) NewAccessToken(cfg *config.Config, customerPassport *customer.Claims) string {

	claims := customer.AuthClaims{
		Claims: &customer.Claims{
			Id: customerPassport.Id,
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

func (r *customerRepository) NewRefreshToken(cfg *config.Config, customerPassport *customer.Claims) string {
	claims := customer.AuthClaims{
		Claims: &customer.Claims{
			Id: customerPassport.Id,
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

func (r *customerRepository) ReloadToken(cfg *config.Config, customerPassport *customer.Claims) string {

	claims := customer.AuthClaims{
		Claims: &customer.Claims{
			Id: customerPassport.Id,
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

	reloadToken, err := token.SignedString([]byte(cfg.Jwt.RefreshSecretKey))
	if err != nil {
		log.Printf("Error: Reoload Token Failed %s", err.Error())
		return err.Error()
	}

	return reloadToken
}

func (r *customerRepository) FindAccessToken(pctx context.Context, accessToken string) (*customer.Credential, error) {
	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	db := r.db.Database("customer_db")
	col := db.Collection("customer_auth")

	result := new(customer.Credential)

	if err := col.FindOne(ctx, bson.M{"access_token": accessToken}).Decode(result); err != nil {
		log.Printf("Error: FindAccessToken Failed %s", err.Error())
		return nil, errors.New("error: email or password invalid")
	}

	return result, nil

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

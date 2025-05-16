package customerRepositories

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/customer"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CustomerRepositoryMock struct {
		mock.Mock
	}
)

func NewCustomerRepoMock() CustomerRepositoryService {
	return &CustomerRepositoryMock{}
}

func (m *CustomerRepositoryMock) FindOneCustomerWithCredential(pctx context.Context, email string) (*customer.Customer, error) {
	return nil, nil
}
func (m *CustomerRepositoryMock) InsertCustomer(pctx context.Context, req *customer.Customer) (primitive.ObjectID, error) {
	return primitive.NilObjectID, nil
}
func (m *CustomerRepositoryMock) FindCustomer(pctx context.Context, customerId string) (*customer.Customer, error) {
	return nil, nil
}
func (m *CustomerRepositoryMock) InsertCustomerCredential(pctx context.Context, req *customer.Credential) (primitive.ObjectID, error) {
	return primitive.NilObjectID, nil
}
func (m *CustomerRepositoryMock) DeleteCustomerCredential(pctx context.Context, credentialId string) (int64, error) {
	return 0, nil
}
func (m *CustomerRepositoryMock) FindAccessToken(pctx context.Context, accessToken string) (*customer.Credential, error) {
	return nil, nil
}
func (m *CustomerRepositoryMock) FindCustomerCredential(pctx context.Context, credentialId string) (*customer.Credential, error) {
	return nil, nil
}
func (m *CustomerRepositoryMock) UpdateCustomerCredential(pctx context.Context, credentialId string, req *customer.UpdateRefreshToken) error {
	return nil
}
func (m *CustomerRepositoryMock) ForgotPassword(pctx context.Context, email string) error {
	return nil
}
func (m *CustomerRepositoryMock) NewAccessToken(cfg *config.Config, customerPassport *customer.Claims) string {
	return ""
}
func (m *CustomerRepositoryMock) NewRefreshToken(cfg *config.Config, customerPassport *customer.Claims) string {
	return ""
}
func (m *CustomerRepositoryMock) ReloadToken(cfg *config.Config, customerPassport *customer.Claims) string {
	return ""
}
func (m *CustomerRepositoryMock) IsUserAlreadyExist(pctx context.Context, username, email string) bool {
	return false
}

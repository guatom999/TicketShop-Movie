package middlewareRepositories

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
)

type (
	IMiddlewareRepositoryService interface {
		AccessTokenSearch(pctx context.Context, cfg *config.Config, accessToken string) error
	}

	middlwareRepository struct {
	}
)

func NewMiddlewareRepository() IMiddlewareRepositoryService {
	return &middlwareRepository{}
}

func (r *middlwareRepository) AccessTokenSearch(pctx context.Context, cfg *config.Config, accessToken string) error {

	_, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	body := bytes.NewBuffer([]byte(nil))

	// Create a new POST request
	req, err := http.NewRequest("POST", fmt.Sprintf("http://"+cfg.AppUrl.CustomerUrl+"/user/find-access-token"), body)
	// req, err := http.NewRequest("POST", "http://localhost:8100/user/find-access-token", body)
	if err != nil {
		log.Printf("Error creating request: %s", err.Error())
		return err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error creating client: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		return errors.New("errors: find access token failed")
	}
	return nil
}

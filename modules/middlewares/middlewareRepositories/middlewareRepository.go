package middlewareRepositories

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type (
	IMiddlewareRepositoryService interface {
		AccessTokenSearch(pctx context.Context, accessToken string) error
	}

	middlwareRepository struct {
	}
)

func NewMiddlewareRepository() IMiddlewareRepositoryService {
	return &middlwareRepository{}
}

func (r *middlwareRepository) AccessTokenSearch(pctx context.Context, accessToken string) error {

	_, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	body := bytes.NewBuffer([]byte(nil))

	// Create a new POST request
	req, err := http.NewRequest("POST", "http://localhost:8100/user/find-access-token", body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	if statusCode != 200 {
		return errors.New("errors: find access token failed")
	}
	return nil
}

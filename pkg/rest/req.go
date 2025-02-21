package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Request(url string) (string, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Error sending GET request: %s", err.Error)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error reading response body: %s", err.Error())
		return "", err
	}

	return string(body), nil
}

func ReqWithParams(url, params string) (any, error) {

	url = fmt.Sprintf(url, params)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Error sending GET request: %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error reading response body: %s", err.Error())
		return "", err
	}

	return string(body), nil
}

func ReqWithBody(url string, body map[string]string) error {
	// Define the URL and body
	bodyData, err := json.Marshal(body)
	if err != nil {
		log.Printf("Error : Failed to Marshall body :%s", err.Error())
		return err
	}

	// Create a new request with method GET and the body
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(bodyData)))
	if err != nil {
		log.Printf("Error creating request: %s", err.Error())
		return err
	}

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	// Read and print the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %s", err.Error())
		return err
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(respBody))

	return nil

}

func Post(url string, value any) error {

	// Define the URL
	// Define the body with access token and other data
	accessToken := "your-access-token"
	requestBody := fmt.Sprintf(`{"accessToken": "%s", "key": "value"}`, value)

	// Convert the body to a byte buffer
	body := bytes.NewBuffer([]byte(requestBody))

	// Create a new POST request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Printf("Error creating request: %s", err.Error())
		return nil
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err.Error())
		return nil
	}
	defer resp.Body.Close()

	// Read and print the response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %s", err.Error())
		return nil
	}

	fmt.Println("Response Status:", resp.Status)

	return nil
}

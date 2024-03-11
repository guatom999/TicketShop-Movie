package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Request(url string) (string, error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return string(body), nil
}

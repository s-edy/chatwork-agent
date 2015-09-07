package chatwork

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (this *Client) Send(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Status: %d", response.StatusCode)
	log.Printf("X-RateLimit-Limit: %s", response.Header.Get("X-RateLimit-Limit"))
	log.Printf("X-RateLimit-Remaining: %s", response.Header.Get("X-RateLimit-Remaining"))
	log.Printf("X-RateLimit-Reset: %s", response.Header.Get("X-RateLimit-Reset"))
	log.Printf("Body: %s", responseBody)

	return responseBody, nil
}

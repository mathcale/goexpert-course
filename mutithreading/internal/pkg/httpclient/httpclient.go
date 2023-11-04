package httpclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HttpClient struct {
	BaseURL string
}

func NewHttpClient(baseURL string) *HttpClient {
	return &HttpClient{
		BaseURL: baseURL,
	}
}

func (c HttpClient) Get(endpoint string, ch chan<- interface{}) error {
	path := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("unexpected status code [%d]", resp.StatusCode))
	}

	defer resp.Body.Close()

	var responseObj interface{}

	if err := json.NewDecoder(resp.Body).Decode(&responseObj); err != nil {
		return err
	}

	ch <- responseObj

	return nil
}

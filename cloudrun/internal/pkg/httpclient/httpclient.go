package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HttpClientInterface interface {
	Get(endpoint string, responseObj interface{}) *HttpClientError
}

type HttpClientError struct {
	Error      error
	StatusCode *int
}

type HttpClient struct {
	BaseURL string
	Timeout time.Duration
}

func NewHttpClient(baseURL string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		BaseURL: baseURL,
		Timeout: timeout,
	}
}

func (c HttpClient) Get(endpoint string, responseObj interface{}) *HttpClientError {
	httpCtx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	path := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	req, err := http.NewRequestWithContext(httpCtx, "GET", path, nil)

	if err != nil {
		return &HttpClientError{
			Error: err,
		}
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &HttpClientError{
			Error:      err,
			StatusCode: &resp.StatusCode,
		}
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseObj); err != nil {
		return &HttpClientError{
			Error:      err,
			StatusCode: &resp.StatusCode,
		}
	}

	return nil
}

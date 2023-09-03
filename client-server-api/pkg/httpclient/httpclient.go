package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HttpClient struct {
	BaseURL string
	Timeout time.Duration
}

func (c HttpClient) Get(endpoint string, responseObj interface{}) error {
	httpCtx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	path := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	req, err := http.NewRequestWithContext(httpCtx, "GET", path, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseObj); err != nil {
		return err
	}

	return nil
}

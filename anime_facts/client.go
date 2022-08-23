package anime_facts

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client  *http.Client
	baseUrl string
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout not provided")
	}
	return &Client{
		baseUrl: BaseApiURL,
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) GetRequest(url string, r interface{}) error {
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return err
	}

	return nil
}

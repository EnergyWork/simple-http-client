package anime_facts

import (
	"errors"
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

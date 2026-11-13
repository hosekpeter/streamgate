package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Option func(*Client)
type Client struct {
	httpClient *http.Client
}

func New(options ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func (c *Client) MakeRequest(ctx context.Context, r Request) (*http.Response, error) {
	target, err := url.JoinPath(r.BaseURL, r.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to join path: %w", err)
	}

	if len(r.Query) > 0 {
		target += "?" + r.Query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, r.Method, target, r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if r.Header != nil {
		req.Header = r.Header.Clone()
	}

	return c.httpClient.Do(req)
}

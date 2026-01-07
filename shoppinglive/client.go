package shoppinglive

import (
	"net/http"
)

type Client struct {
	client *http.Client
	header map[string]string
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		client: client,
		header: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120 Safari/537.36",
		},
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	for key, value := range c.header {
		req.Header.Set(key, value)
	}

	return c.client.Do(req)
}

func (c *Client) GetServiceType() string {
	return "SHOPPING"
}

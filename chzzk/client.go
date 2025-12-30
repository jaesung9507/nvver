package chzzk

import (
	"net/http"

	"github.com/google/uuid"
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
			"User-Agent":                 "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120 Safari/537.36",
			"Origin":                     "https://chzzk.naver.com",
			"Deviceid":                   uuid.New().String(),
			"Front-Client-Platform-Type": "PC",
			"Front-Client-Product-Type":  "web",
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
	return "CHZZK"
}

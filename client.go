package nvver

import "net/http"

type Client interface {
	GetServiceType() string
	Do(req *http.Request) (*http.Response, error)
}

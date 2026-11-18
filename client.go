package nvver

import (
	"errors"
	"net/http"
)

type Client interface {
	GetServiceType() string
	Do(req *http.Request) (*http.Response, error)
}

var ErrServiceDiscontinued = errors.New("service has been discontinued")

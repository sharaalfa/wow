package httpclient

import "net/http"

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type RealHTTPClient struct{}

func (c *RealHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

package gohttp

import (
	"net/http"
)

type httpClient struct {
	builder *clientBuilder
	client  *http.Client
}

type Client interface {

	// Methods for interacting with HTTP calls
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	response, err := c.do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, nil)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

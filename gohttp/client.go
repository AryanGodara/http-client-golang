package gohttp

import (
	"net/http"
	"sync"
)

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

type Client interface {

	// Methods for interacting with HTTP calls
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*Response, error)
	Patch(url string, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Options(url string, headers ...http.Header) (*Response, error)
}

func (c *httpClient) Get(url string, headers ...http.Header) (*Response, error) {
	response, err := c.do(http.MethodGet, url, getHeaders(headers...), nil)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), body)
}

func (c *httpClient) Patch(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), nil)
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}

func (c *httpClient) Options(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}

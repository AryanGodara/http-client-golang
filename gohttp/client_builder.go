package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
}

type ClientBuilder interface {
	// Methods for configuration
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder

	Build() Client
}

func (c *clientBuilder) Build() Client {
	client := &httpClient{
		headers:            c.headers,
		maxIdleConnections: c.maxIdleConnections,
		connectionTimeout:  c.connectionTimeout,
		responseTimeout:    c.responseTimeout,
		disableTimeouts:    c.disableTimeouts,
	}
	return client
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}
func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}
func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnections = connections
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

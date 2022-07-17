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
	client             *http.Client
	userAgent          string
}

type ClientBuilder interface {
	// Methods for configuration
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	SetUserAgent(u string) ClientBuilder

	Build() Client
}

// Build creates a new client from the clientBuilder
func (c *clientBuilder) Build() Client {
	client := &httpClient{
		builder: c,
	}

	return client
}

// Newbuilder creates a new ClientBuilder, used to create custom clients
func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}

	return builder
}

// SetHeaders sets the Headers for the client builder
func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers

	return c
}

// SetConnectionTimeout sets the Connection Timeout for the client builder
func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout

	return c
}

// SetConnectionTimeout sets the Response Timeout for the client builder
func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout

	return c
}

// SetConnectionTimeout sets the Max Idle Connections for the client builder
func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnections = connections

	return c
}

// SetConnectionTimeout disables the Timeouts for the client builder
func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable

	return c
}

// SetHttpClient sets the http.Client for the client builder
func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client

	return c
}

// SetHttpClient sets the userAgent for the client builder
func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent

	return c
}

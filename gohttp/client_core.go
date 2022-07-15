package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"time"
)

const (
	defaultMaxIdleTimeout    = 5
	defaultConnectionTimeout = 1 * time.Second
	defaultResponseTimeout   = 5 * time.Second
)

func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	c.client = &http.Client{
		Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
			ResponseHeaderTimeout: c.getResponseTimeout(),
			DialContext: net.Dialer{
				Timeout: c.getConnectionTimeout(),
			}.Resolver.Dial,
		},
	}

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	} else {
		return defaultMaxIdleTimeout
	}
}
func (c *httpClient) getResponseTimeout() time.Duration {
	if c.responseTimeout > 0 {
		return c.getResponseTimeout()
	} else if c.disableTimeouts {
		return 0
	} else {
		return defaultResponseTimeout
	}

}
func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	} else if c.disableTimeouts {
		return 0
	} else {
		return defaultConnectionTimeout
	}
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch contentType {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body) // json is the default format for body
	}

}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add default headers to the request
	for header, value := range c.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers to the request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	// client := http.Client{}

	fullHeaders := c.getRequestHeaders(headers)
	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		panic(err)
	}

	// create a new http request
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	// return c.client.Do(request)	//? This'll fail, if we run this before creating a client

	client := c.getHttpClient()
	return client.Do(request)
}

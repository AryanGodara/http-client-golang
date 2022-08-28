package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/AryanGodara/http-client-golang/core"
	"github.com/AryanGodara/http-client-golang/gohttp_mock"
)

const (
	defaultMaxIdleTimeout    = 5
	defaultConnectionTimeout = 1 * time.Second
	defaultResponseTimeout   = 5 * time.Second
)

func (c *httpClient) getHttpClient() core.HttpClient {
	if gohttp_mock.MockupServer.IsEnabled() {
		return gohttp_mock.MockupServer.GetMockedClient()
	}

	c.clientOnce.Do(func() { // func() is executed only once, even in concurrent enviornments

		if c.builder.client != nil { // someone defined a custom http.Client already, so we need to use that
			c.client = c.builder.client
			return
		}

		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	} else {
		return defaultMaxIdleTimeout
	}
}
func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	} else if c.builder.disableTimeouts {
		return 0
	} else {
		return defaultResponseTimeout
	}

}
func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	} else if c.builder.disableTimeouts {
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

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*core.Response, error) {
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

	//* When mock_server=ON, getHttpClient returns mockhttpclient, which has a custom
	// Do function, which handles mock requests
	response, err := c.getHttpClient().Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close() // So the user of this package doesn't have to do this :)

	finalResponse := core.Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       responseBody,
	}

	return &finalResponse, nil
}

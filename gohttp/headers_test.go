package gohttp

import (
	"net/http"
	"testing"

	"github.com/AryanGodara/http-client-golang/gomime"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	commonHeaders.Set(gomime.HeaderUserAgent, "mocked-http-client")
	client.builder = &clientBuilder{
		headers: commonHeaders,

		// userAgent: "cool-user-agent",
	}

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}

	if finalHeaders.Get(gomime.HeaderContentType) != "application/json" {
		t.Error("invalid content type received")
	}

	if finalHeaders.Get(gomime.HeaderUserAgent) != "mocked-http-client" {
		t.Error("invalid user agent received")
	}
}

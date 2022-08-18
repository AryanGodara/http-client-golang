package gohttp

import (
	"net/http"
	"testing"

	"github.com/AryanGodara/http-client-golang/gomime"
)

func TestGetHeaders(t *testing.T) {
	t.Run("no headers", func(t *testing.T) {
		ret_header := getHeaders()
		if len(ret_header) != 0 {
			t.Error("expected an empty http Header")
		}
	})

	t.Run("one header", func(t *testing.T) {
		header := make(http.Header)
		header.Set("ABC", "XYZ")
		ret_header := getHeaders(header)

		if len(ret_header) == 0 {
			t.Error("exptected to get a non nil http Header")
		}
		if ret_header.Get("ABC") != "XYZ" {
			t.Error("mismathed key-value pair")
		}
	})

	t.Run("multiple headers", func(t *testing.T) {
		header := make(http.Header)
		header.Set("ABC", "XYZ")

		extra_header := make(http.Header)
		extra_header.Set("RAN", "DOM")

		ret_header := getHeaders(header, extra_header)
		if len(ret_header) == 0 {
			t.Error("exptected to get a non nil http Header")
		}
		if len(ret_header) > 1 {
			t.Error("exptected to get only a single http Header")
		}
		if ret_header.Get("ABC") != "XYZ" {
			t.Error("mismathed key-value pair")
		}
	})
}

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

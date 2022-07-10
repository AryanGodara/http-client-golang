package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//? Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	//? Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-request-id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	//? Validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers")
	}
}

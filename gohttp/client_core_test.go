package gohttp

import (
	"encoding/xml"
	"fmt"
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

	if finalHeaders.Get("X-Request-id") != "ABC-123" {
		t.Error("invalid request id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent received")
	}
}

func TestGetRequestBody(t *testing.T) {

	//? Initialization
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		//? Execution
		body, err := client.getRequestBody("", nil)

		//? Validation
		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("no body expected when passing a nil body")
		}

	})

	t.Run("BodyWithJson", func(t *testing.T) {

		//? Execution
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		//? Validation
		if err != nil {
			t.Error("no error expected when passing a json body")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}

	})

	t.Run("BodyWithXml", func(t *testing.T) {

		//? Execution
		type User struct {
			FirstName string `xml:"name>first"`
			LastName  string `xml:"name>last"`
		}

		v := &User{FirstName: "John", LastName: "Doe"}
		requestBody, _ := xml.Marshal(v)
		body, err := client.getRequestBody("application/xml", requestBody)

		//? Validation
		if err != nil {
			fmt.Println(err.Error())
			t.Error("no error expected when passing an xml body")
		}

		xmlbody := `[60 85 115 101 114 62 60 70 105 114 115 116 78 97 109 101 62 74 111 104 110 60 47 70 105 114 115 116 78 97 109 101 62 60 76 97 115 116 78 97 109 101 62 68 111 101 60 47 76 97 115 116 78 97 109 101 62 60 47 85 115 101 114 62]`

		if string(body) != xmlbody {
			t.Error("invalid xml body obtained")
		}

	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {

		//? Execution
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("", requestBody)

		//? Validation
		if err != nil {
			t.Error("no error expected when passing a json body")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}

	})
}

//! Using sub-tests instead of creating a new test for each return statement in a function
/*
func TestGetRequestNilBody(t *testing.T) {
	//? Initializaiton
	client := httpClient{}

	//? Execution
	body, err := client.getRequestBody("", nil)

	//? Validation
	if err != nil {
		t.Error("no error expected when passing a nil body")
	}

	if body != nil {
		t.Error("no body expected when passing a nil body")
	}
}
func TestGetRequestJsnoBody(t *testing.T) {
	//? Initializaiton
	client := httpClient{}

	//? Execution

	//? Validation
}
func TestGetRequestXmlBody(t *testing.T) {
	//? Initializaiton
	client := httpClient{}

	//? Execution

	//? Validation
}
func TestGetRequestJsonAsDefaultBody(t *testing.T) {
	//? Initializaiton
	client := httpClient{}

	//? Execution

	//? Validation
}
*/

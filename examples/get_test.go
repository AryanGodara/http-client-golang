package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/AryanGodara/http-client-golang/gohttp"
)

func TestGetEndpoint(t *testing.T) {
	// Tell the HTTP library to mock any further requests from here
	gohttp.StartMockServer()

	GetEndpoints()

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// Initialization:
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initialization:
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`, // returns 123, instead of valid string
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "json unmarshal error" {
			t.Error("invalid error message received")
		}
	})
	t.Run("TestNoError", func(t *testing.T) {
		// Initialization:
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation
		if err != nil {
			t.Errorf("no error was expected, and we got: %s\n", err.Error())
		}

		if endpoints == nil {
			t.Error("endpoints were expected, but we got nil")
		}

		if endpoints != nil && endpoints.RepositoryUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}

	})

}

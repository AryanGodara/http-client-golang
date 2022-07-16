package gohttp

import (
	"fmt"
	"net/http"
)

type Mock struct {
	Method             string
	Url                string
	RequestBody        string
	ResponseBody       string
	ResponseStatusCode int
	Error              error
}

func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body:       []byte(m.ResponseBody),
	}

	return &response, nil
}

package examples

import "fmt"

/*
	"current_user_url": "https://api.github.com/user",
	"authorizations_url": "https://api.github.com/authorization",
	"repository_url": "https://api.github.com/repos/{owner}/{repo}"
*/

type Endpoints struct {
	CurrentUser       string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("1. Status: ", response.Status())
	fmt.Println("2. StatusCode: ", response.StatusCode())
	fmt.Println("3. ResponseBody: ", response.StringBody())

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println("Current User: ", endpoints.CurrentUser)

	return &endpoints, nil
}

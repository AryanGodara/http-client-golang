package examples

import "fmt"

/*
	"current_user_url": "https://api.github.com/user",
	"authorizations_url": "https://api.github.com/authorization",
	"repository_url": "https://api.github.com/repos/{owner}/{repo}"
*/

type Endpoints struct {
	CurrentUser       string `'json:"current_user_url"`
	AuthorizationsUrl string `'json:"authorizations_url"`
	RepositoryUrl     string `'json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.StringBody())

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Printf("Repositories URL: %s\n", endpoints.RepositoryUrl)

	return &endpoints, nil
}

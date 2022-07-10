package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AryanGodara/http-client-golang/gohttp"
)

var (
	githubHttpClient = getGitHubClient()
)

func getGitHubClient() gohttp.HttpClient {
	client := gohttp.New()

	// Creates a map[string][]string (key:string, val: slice of string)
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client.SetHeaders(commonHeaders)

	return client
}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getUrls() {

	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func main() {
	// getUrls()

	v := &User{FirstName: "John", LastName: "Doe"}
	xmlbody, _ := xml.Marshal(v)

	fmt.Println(xmlbody)
}

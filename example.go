package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/AryanGodara/http-client-golang/gohttp"
)

var (
	githubHttpClient = getGitHubClient()
)

func getGitHubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(false).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(50 * time.Millisecond).
		SetMaxIdleConnections(5). // all return gohttp.ClientBuilder
		Build()                   // return gohttp.Client

	return client
}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	getUrls()

	// v := &User{FirstName: "John", LastName: "Doe"}
	// xmlbody, _ := xml.Marshal(v)

	// fmt.Println(string(xmlbody))
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

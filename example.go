package main

import (
	"fmt"
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

	//* Using our custom response
	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())

	var user User
	err = response.UnmarshalJson(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StringBody())

	//? Using default http.response
	/*
		defer response.Body.Close()
		fmt.Println(response.StatusCode)

		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(bytes))

		var user User
		if err := json.Unmarshal(bytes, &user); err != nil {
			panic(err)
		}

		fmt.Println(user.FirstName, user.LastName)
	*/
}

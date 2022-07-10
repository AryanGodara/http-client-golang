package go_httpclient

import (
	"fmt"

	"github.com/AryanGodara/go-httpclient/gohttp"
)

func basicExample() {
	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}

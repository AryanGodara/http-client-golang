package examples

import (
	"time"

	"github.com/AryanGodara/http-client-golang/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()

	return client
}

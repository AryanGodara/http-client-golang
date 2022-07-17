package examples

import (
	"net/http"
	"time"

	"github.com/AryanGodara/http-client-golang/gohttp"
	"github.com/AryanGodara/http-client-golang/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Aryan-Computer").
		Build()

	return client
}

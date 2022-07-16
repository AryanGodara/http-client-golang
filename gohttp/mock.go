package gohttp

type Mock struct {
	Method             string
	Url                string
	RequestBody        string
	ResponseBody       string
	ResponseStatusCode int
	Error              error
}

package builder

import (
	"fmt"
	"net/url"
)

type Request struct {
	Method      string
	URL         string
	Headers     map[string]string
	QueryParams url.Values
	Body        string
}

func (r Request) Describe() {
	fmt.Printf("[HTTP Request]\n")
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL)
	fmt.Printf("Headers: %v\n", r.Headers)
	fmt.Printf("Query Params: %v\n", r.QueryParams.Encode())
	fmt.Printf("Body: %s\n", r.Body)
}

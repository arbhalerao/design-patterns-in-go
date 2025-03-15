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

// Builder Interface
type RequestBuilder interface {
	SetMethod(method string) RequestBuilder
	SetURL(url string) RequestBuilder
	SetHeader(key, value string) RequestBuilder
	SetQueryParam(key, value string) RequestBuilder
	SetBody(body string) RequestBuilder
	Build() Request
}

type HttpRequestBuilder struct {
	request Request
}

func NewRequestBuilder() *HttpRequestBuilder {
	return &HttpRequestBuilder{
		request: Request{
			Headers:     make(map[string]string),
			QueryParams: url.Values{},
		},
	}
}

func (b *HttpRequestBuilder) SetMethod(method string) RequestBuilder {
	b.request.Method = method
	return b
}

func (b *HttpRequestBuilder) SetURL(u string) RequestBuilder {
	b.request.URL = u
	return b
}

func (b *HttpRequestBuilder) SetHeader(key, value string) RequestBuilder {
	b.request.Headers[key] = value
	return b
}

func (b *HttpRequestBuilder) SetQueryParam(key, value string) RequestBuilder {
	b.request.QueryParams.Add(key, value)
	return b
}

func (b *HttpRequestBuilder) SetBody(body string) RequestBuilder {
	b.request.Body = body
	return b
}

func (b *HttpRequestBuilder) Build() Request {
	return b.request
}

func (r Request) Describe() {
	fmt.Printf("\n[HTTP Request]\n")
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL)
	fmt.Printf("Headers: %v\n", r.Headers)
	fmt.Printf("Query Params: %v\n", r.QueryParams.Encode())
	fmt.Printf("Body: %s\n", r.Body)
}

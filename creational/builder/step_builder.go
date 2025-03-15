package builder

import (
	"net/url"
)

type MethodStep interface {
	SetMethod(method string) URLStep
}

type URLStep interface {
	SetURL(u string) HeaderStep
}

type HeaderStep interface {
	SetHeader(key, value string) HeaderStep
	SetQueryParam(key, value string) HeaderStep
	SetBody(body string) BuildStep
}

type BuildStep interface {
	Build() Request
}

type HttpRequestStepBuilder struct {
	request Request
}

func NewRequestStepBuilder() MethodStep {
	return &HttpRequestStepBuilder{
		request: Request{
			Headers:     make(map[string]string),
			QueryParams: url.Values{},
		},
	}
}

func (sb *HttpRequestStepBuilder) SetMethod(method string) URLStep {
	sb.request.Method = method
	return sb
}

func (sb *HttpRequestStepBuilder) SetURL(u string) HeaderStep {
	sb.request.URL = u
	return sb
}

func (sb *HttpRequestStepBuilder) SetHeader(key, value string) HeaderStep {
	sb.request.Headers[key] = value
	return sb
}

func (sb *HttpRequestStepBuilder) SetQueryParam(key, value string) HeaderStep {
	sb.request.QueryParams.Add(key, value)
	return sb
}

func (sb *HttpRequestStepBuilder) SetBody(body string) BuildStep {
	sb.request.Body = body
	return sb
}

func (sb *HttpRequestStepBuilder) Build() Request {
	return sb.request
}

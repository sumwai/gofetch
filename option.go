package gofetch

import (
	"encoding/json"
	"strings"
	"time"
)

type (
	Option func(req *Request)
)

// WithUrl set request url
func WithUrl(url string) Option {
	return func(req *Request) {
		req.Url = url
	}
}

// WithParam set query params item
func WithParam(name, value any) Option {
	return func(req *Request) {
		req.Params[name] = value
	}
}

// WithParams set query params
func WithParams(params Params) Option {
	return func(req *Request) {
		req.Params = params
	}
}

// WithForm set form data when method is POST
func WithForm(data Params) Option {
	return func(req *Request) {
		req.With(WithHeader("Content-Type", "application/x-www-form-urlencoded"))
		req.BodyReader = strings.NewReader(data.String())
	}
}

// WithJson set json data when method is POST
func WithJson(data any) Option {
	return func(req *Request) {
		jsondata, err := json.Marshal(data)
		if err != nil {
			return
		}
		req.With(WithHeader("Content-Type", "application/json"))
		req.BodyReader = strings.NewReader(string(jsondata))
	}
}

// WithBody set body data when method is POST
func WithBody(body string) Option {
	return func(req *Request) {
		req.With(WithHeader("Content-Type", "application/x-www-form-urlencoded"))
		req.BodyReader = strings.NewReader(body)
	}
}

// WithHeader set request headers item
func WithHeader(key, value string) Option {
	return func(req *Request) {
		req.Headers[key] = value
	}
}

// WithHeaders set request headers
func WithHeaders(headers Headers) Option {
	return func(req *Request) {
		req.Headers = headers
	}
}

// WithMethod set request method
func WithMethod(method string) Option {
	return func(req *Request) {
		req.Method = method
	}
}

// WithTimeout set request timeout
func WithTimeout(timeout time.Duration) Option {
	return func(req *Request) {
		req.Timeout = timeout
	}
}

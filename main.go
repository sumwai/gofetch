package gofetch

import (
	"time"
)

func New() *Request {
	return &Request{
		Params:  Params{},
		Timeout: time.Second * 10,
		Headers: map[string]string{},
		Method:  "GET",
	}
}

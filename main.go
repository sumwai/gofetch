package gofetch

import (
	"log"
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

func main() {
	client := New()
	client.With(WithHeader("Content-Type", "application/json"))
	client.With()
	ret, err := client.GET(
		"http://47.243.43.76",
		nil,
	)
	if err != nil {
		panic(err)
	}
	log.Println(string(ret))
}

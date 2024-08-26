package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

type (
	Headers map[string]string
	Request struct {
		Method     string
		Url        string
		Headers    Headers
		Params     Params
		Timeout    time.Duration
		BodyReader io.Reader
	}
)

func (r *Request) With(options ...Option) *Request {
	for _, option := range options {
		option(r)
	}
	return r
}

func (r *Request) Fetch(options ...Option) (body []byte, err error) {
	return r.With(options...).do()
}

// POST simple post request with form data
func (r *Request) POST(url string, data Params) (body []byte, err error) {
	return r.With(
		WithMethod("POST"),
		WithUrl(url),
		WithForm(data),
	).do()
}

// GET simple get request with query params
func (r *Request) GET(url string, query Params) (body []byte, err error) {
	return r.With(
		WithMethod("GET"),
		WithUrl(url),
		WithParams(query),
	).do()
}

// POSTJson post request with json body
func (r *Request) POSTJson(url string, data any) (body []byte, err error) {
	return r.With(
		WithMethod("POST"),
		WithUrl(url),
		WithJson(data),
	).do()
}

func (r *Request) do() (body []byte, err error) {
	client := http.Client{
		Timeout: r.Timeout,
	}
	if r.Params.String() != "" {
		r.Url = r.Url + "?" + r.Params.String()
	}
	// b, _ := io.ReadAll(r.BodyReader)
	// log.Println(string(b))
	req, err := http.NewRequest(r.Method, r.Url, r.BodyReader)
	if err != nil {
		return
	}
	// set headers
	for name, value := range r.Headers {
		req.Header.Set(name, value)
	}
	log.Println(r)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

package router

import (
	"bytes"
	"net/http"
)

type response struct {
	h http.Header
	b bytes.Buffer
	s int
}

func newResponse() response {
	return response{
		h: http.Header{},
	}
}

func (r *response) Header() http.Header {
	return r.h
}

func (r *response) Write(b []byte) (int, error) {
	return r.b.Write(b)
}

func (r *response) WriteHeader(statusCode int) {
	r.s = statusCode
}

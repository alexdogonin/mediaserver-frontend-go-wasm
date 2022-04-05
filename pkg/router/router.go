package router

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	r chi.Router
}

func NewRouter() Router {
	return Router{
		r: chi.NewRouter(),
	}
}

func (r *Router) Handle(pattern string, fn func()) {
	r.r.Get(pattern, func(_ http.ResponseWriter, _ *http.Request) {
		fn()
	})
}

func (r *Router) Route(pattern string, fn func(router *Router)) Router {

	return Router{r.r.Route(pattern, func(rr chi.Router) {
		fn(&Router{r: rr})
	})}
}

func (r *Router) Serve(href string) error {
	u, err := url.Parse(href)
	if err != nil {
		return err
	}

	req := http.Request{
		Method: http.MethodGet,
		URL:    u,
	}

	resp := newResponse()

	r.r.ServeHTTP(&resp, &req)

	if resp.s != http.StatusOK {
		return errors.New(resp.b.String())
	}

	return nil
}

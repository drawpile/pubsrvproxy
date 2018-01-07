package main

import (
	"github.com/drawpile/pubsrvproxy/queries"
	"net/http"
	"net/url"
	"strings"
)

// An extended ServeMux for the API endpoints
type apiMux struct {
	*http.ServeMux
	cfg   *config
	cache *queries.QueryCache
}

// API request context
type apiContext struct {
	path  string
	query url.Values
	cfg   *config
	cache *queries.QueryCache
}

func (ctx *apiContext) GetQueryOpts() queries.QueryOpts {
	return queries.QueryOpts{
		ServerAddr: ctx.cfg.ServerAddr,
		Cache:      ctx.cache,
	}
}

func (mux *apiMux) HandleApiEndpoint(prefix string, endPoint func(*apiContext) apiResponse) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			methodNotAllowedResponse{"GET"}.WriteResponse(w)
		}
		if p := strings.TrimPrefix(req.URL.Path, prefix); len(p) < len(req.URL.Path) {
			// Build request context
			ctx := apiContext{
				path:  p,
				query: url.Values{},
				cfg:   mux.cfg,
				cache: mux.cache,
			}

			// Call API endpoint
			endPoint(&ctx).WriteResponse(w)

		} else {
			notFoundResponse().WriteResponse(w)
		}
	})
}

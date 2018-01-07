package main

import (
	"github.com/drawpile/pubsrvproxy/queries"
	"log"
	"net/http"
)

func StartServer(cfg *config) {
	mux := &apiMux{http.NewServeMux(), cfg, &queries.QueryCache{ExpirationTime: cfg.CacheTime}}

	mux.HandleApiEndpoint("/sessions/", SessionListEndpoint)

	if cfg.UserView {
		mux.HandleApiEndpoint("/users/", UserListEndpoint)
	}

	log.Println("Starting dpserver public statistics proxy server at", cfg.Listen, "proxying", cfg.ServerAddr)
	log.Fatal(http.ListenAndServe(cfg.Listen, mux))
}

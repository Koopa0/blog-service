package main

import (
	"github.com/koopa0/blog-service/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":8081",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

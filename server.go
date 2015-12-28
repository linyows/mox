package main

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
	mux  *http.ServeMux
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func NewServer() *Server {
	serveMux := http.NewServeMux()

	httpServer := &http.Server{
		Addr:           "localhost:8080",
		Handler:        serveMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s := &Server{
		http: httpServer,
		mux:  serveMux,
	}

	s.mux.HandleFunc("/", handler)

	return s
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

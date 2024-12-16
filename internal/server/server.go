package server

import (
	"net/http"
	"time"
)

func NewServer() *http.Server {

	server := &http.Server{
		Addr:         "localhost:5050",
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server

}

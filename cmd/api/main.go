package main

import (
	"errors"
	"fmt"
	"github.com/kirillshch/test_go_effective/config"
	"io"
	"log"
	"net/http"
)

func main() {
	c := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Printf("Starting server " + s.Addr)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Hello world!")

	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

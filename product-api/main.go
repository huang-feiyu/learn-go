package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"huangblog.com/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	// sm mux. it is created by default, but we can create our own.
	sm := http.NewServeMux()
	// register the handler
	sm.Handle("/", hh)
	sm.Handle("/bye", gb)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received signal:", sig, "Shutting down...")

	// everything is done, so we can exit
	ct, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ct)
}
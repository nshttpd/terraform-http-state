package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"terraform-http-state/handlers"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	StatePath = "/state"
)

func main() {
	logLevel := flag.String("log-level", "info", "log level")
	logFormat := flag.String("log-format", "json", "log format text or json (default json)")
	statePath := flag.String("state-path", StatePath, "base path for state storage")
	port := flag.String("port", ":8080", "port to listen on")

	flag.Parse()

	ll, err := log.ParseLevel(*logLevel)

	if err != nil {
		log.Fatal("unknown log level : ", *logLevel)
	}
	log.SetLevel(ll)
	if *logFormat == "text" {
		log.SetFormatter(&log.TextFormatter{})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}

	r := mux.NewRouter()
	s := r.PathPrefix(*statePath).Subrouter()
	s.HandleFunc("/{key}/{path:.*}", handlers.GetHandler).Methods("GET")
	r.HandleFunc("/healthz", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		//noinspection GoUnhandledErrorResult
		fmt.Fprintf(res, "OK")
	})

	srv := &http.Server{
		Addr:         *port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.WithField("port", *port).Info("starting server")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	//noinspection GoUnhandledErrorResult
	srv.Shutdown(ctx)
	log.Info("shutting down")
	os.Exit(0)
}

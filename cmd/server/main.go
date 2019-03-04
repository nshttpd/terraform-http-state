package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

func main() {
	logLevel := flag.String("log-level", "info", "log level")
	logFormat := flag.String("log-format", "json", "log format text or json (default json)")
	//	port := flag.String("port", "8080", "port to listen on")

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

}

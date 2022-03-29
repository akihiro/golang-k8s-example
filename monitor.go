package main

import (
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Monitor string
)

func init() {
	flag.StringVar(&Monitor, "monitor", ":12000", "listen address for monitor and profile")
}

func SetupMonitor() {
	if Monitor == "" {
		return
	}
	http.Handle("/metrics", promhttp.Handler())
	srv := http.Server{
		Addr: Monitor,
	}
	go func() {
		log.Printf("Start monitoring: %s", Monitor)
		err := srv.ListenAndServe()
		var opErr *net.OpError
		if errors.As(err, &opErr) {
			log.Fatal(opErr)
		}
		log.Print(err)
	}()
}

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	Listen       string
	ShutdownWait time.Duration
)

func init() {
	flag.StringVar(&Listen, "listen", ":2000", "Listen address")
	flag.DurationVar(&ShutdownWait, "shutdown", time.Second*5, "Wait timer")
}

func main() {
	flag.VisitAll(func(f *flag.Flag) {
		key := strings.ToUpper(f.Name)
		if val, ok := os.LookupEnv(key); ok {
			f.Value.Set(val)
		}
	})
	flag.Parse()
	SetupMonitor()
	mux := http.NewServeMux()
	srv := http.Server{
		Addr:    Listen,
		Handler: mux,
	}
	go func() {
		log.Print(srv.ListenAndServe())
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	sig := <-sigCh
	log.Printf("Recieve signal: %s", sig)

	ctx, stop := context.WithTimeout(context.Background(), ShutdownWait)
	defer stop()
	if err := srv.Shutdown(ctx); err != nil {
		log.Print(err)
	}
	log.Print("finished")
}

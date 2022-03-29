package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	srv := http.Server{
		Addr: ":2000",
	}
	go func() {
		log.Print(srv.ListenAndServe())
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	sig := <-sigCh
	log.Printf("Recieve signal: %s", sig)

	ctx, stop := context.WithTimeout(context.Background(), time.Second*5)
	defer stop()
	if err := srv.Shutdown(ctx); err != nil {
		log.Print(err)
	}
	log.Print("finished")
}

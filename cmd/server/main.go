package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Info().Msg("Starting server")

	r := NewMuxHandler()
	port := 8080

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Panic().
				Err(err).
				Int("port", port).
				Msgf("cannot start server on port %d", port)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Info().Msg("Closing server...")
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	os.Exit(0)
}
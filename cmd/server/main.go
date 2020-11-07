package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Info().Msg("Starting server")

	r := NewMuxHandler()
	port := 8080

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic().
				Err(err).
				Int("port", port).
				Msgf("cannot start server on port %d", port)
		}
	}()
	log.Info().Int("port", port).Msgf("Server running on port %d", port)
	<-done

	log.Info().Msg("Closing server...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msg("Server closing failed")
	}
	log.Info().Msg("Server closed")

	os.Exit(0)
}

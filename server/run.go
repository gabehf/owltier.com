package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mnrva-dev/owltier.com/server/auth"
	"github.com/mnrva-dev/owltier.com/server/config"
)

// boilerplate taken from go-chi on GitHub
func Run() {
	// The HTTP Server
	server := &http.Server{Addr: config.ListenAddr(), Handler: handler()}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		log.Println("* Server interrupted, shutting down...")
		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, timeoutCtx := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("* Graceful shutdown timed out, forcing exit...")
			}
		}()

		/*
		*	Graceful shutdown logic goes here
		 */

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		timeoutCtx()
		serverStopCtx()
	}()

	// Run the server
	log.Println("* Listening on http://" + config.ListenAddr())
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	log.Println("* Shutdown complete.")
}

func handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/auth", auth.BuildRouter())

	return r
}

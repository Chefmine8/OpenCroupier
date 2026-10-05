package api

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Chefmine8/OpenCroupier/internal/env"
)

func Listen(server *http.Server) {
	port := env.GetEnvVariable("PORT", "42010")
	ip := env.GetEnvVariable("HOST_IP", "127.0.0.1")
	log.Printf("HTTPS server start at https://%s:%s\n", ip, port)
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Erreur serveur : %v", err)
	}
}

func GestionStop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Force stop issue : %v", err)
	}

	log.Println("Server properly stopped")
}

func NewStop() chan os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	return stop
}

func NewServer(db *sql.DB) *http.Server {
	r := NewApi(db)

	tlsConfig := &tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	port := env.GetEnvVariable("PORT", "42010")
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		TLSConfig:    tlsConfig,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server
}

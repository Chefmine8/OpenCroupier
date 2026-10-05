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
)

func Listen(server *http.Server) {
	log.Println("Serveur HTTPS lancé sur https://localhost:8443")
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Erreur serveur : %v", err)
	}
}

func GestionStop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Échec de l'arrêt forcé : %v", err)
	}

	log.Println("Serveur arrêté proprement.")
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

	server := &http.Server{
		Addr:         ":8443",
		Handler:      r,
		TLSConfig:    tlsConfig,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server
}

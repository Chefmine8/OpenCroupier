package main

import (
	"log"

	"github.com/Chefmine8/OpenCroupier/internal/api"
	sqlite "github.com/Chefmine8/OpenCroupier/internal/sql"
)

func main() {
	server := api.NewServer()
	stop := api.NewStop()
	db := sqlite.OpenDB()
	defer db.Close()

	go func() {
		api.Listen(server)
	}()

	<-stop
	log.Println("Arrêt du serveur en cours...")
	api.GestionStop(server)
}

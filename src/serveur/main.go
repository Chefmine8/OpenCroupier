package main

import (
	"log"

	"github.com/Chefmine8/OpenCroupier/internal/api"
	sqlite "github.com/Chefmine8/OpenCroupier/internal/sql"
)

func main() {
	db := sqlite.OpenDB()
	defer db.Close()
	server := api.NewServer(db)
	stop := api.NewStop()

	go func() {
		api.Listen(server)
	}()

	<-stop
	log.Println("Arrêt du serveur en cours...")
	api.GestionStop(server)
}

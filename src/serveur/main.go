package main

import (
	"log"

	"github.com/Chefmine8/OpenCroupier/internal/api"
	"github.com/Chefmine8/OpenCroupier/internal/env"
	sqlite "github.com/Chefmine8/OpenCroupier/internal/sql"
)

func main() {
	env.NewEnv()

	db := sqlite.OpenDB()
	defer db.Close()
	server := api.NewServer(db)
	stop := api.NewStop()

	go func() {
		api.Listen(server)
	}()

	<-stop
	log.Println("Server stopping")
	api.GestionStop(server)
}

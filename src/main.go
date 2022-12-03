package main

import (
	"fmt"

	"github.com/drumer2142/microWeb/src/api"
	"github.com/drumer2142/microWeb/src/api/database"
	"github.com/drumer2142/microWeb/src/config"
	"github.com/drumer2142/microWeb/src/migrations"
)

func init() {
	config.Load()
	migrations.Load()
}

func main() {
	store, err := database.Connect()

	if err != nil {
		return
	}

	server := api.NewApiServer(fmt.Sprintf(":%d", config.APPPORT), store)
	server.Run()
}

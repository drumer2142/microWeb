package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drumer2142/microWeb/src/api/database"
	"github.com/drumer2142/microWeb/src/api/router"
	"github.com/drumer2142/microWeb/src/config"
	"github.com/drumer2142/microWeb/src/migrations"
)

func init() {
	config.Load()
}

func Run() {
	if err := database.Init(); err != nil {
		log.Fatalf("database: %v", err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("database close: %v", err)
		}
	}()

	migrations.Run()

	port := config.APPPORT
	r := router.New()

	fmt.Printf("\nListening [::]:%d \n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

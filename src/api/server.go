package api

import (
  "fmt"
  "log"
  "net/http"
  "github.com/drumer2142/microWeb/src/config"
  "github.com/drumer2142/microWeb/src/migrations"
  "github.com/drumer2142/microWeb/src/api/router"
)

func init(){
    config.Load()
    migrations.Load()
}

func Run() {
    port := config.APPPORT
    router := router.New()

    fmt.Printf("\nListening [::]:%d \n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

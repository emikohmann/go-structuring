package main

import (
    "github.com/emikohmann/go-structuring/domain-driven/domain/ban"
    "github.com/emikohmann/go-structuring/domain-driven/http"
    "github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
    "github.com/emikohmann/go-structuring/domain-driven/repository/ds"
    "github.com/emikohmann/go-structuring/domain-driven-full/repository/db"
)

func main() {
    //repo := ds.New()
    repo := db.New()
    serv := ban.NewService(repo)

    handler := http.NewHandler(serv)
    router := mlhandlers.DefaultMeliRouter()

    router.POST("/ban", handler.ApplyBAN)
    router.DELETE("/ban", handler.RemoveBAN)

    router.Run()
}

package main

import (
	"log"
	"net/http"

	"github.com/dnlaxe/linkwall/backend/internal/server"
)

func main() {
	addr := ":8080"
	log.Printf("lisiting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.NewRouter()))
}

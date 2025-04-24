package main

import (
	"log"
	"net/http"

	"github.com/spatxos/HotSearch/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/news", api.GetHotListHandler)
	mux.HandleFunc("/api/sources", api.GetSourcesHandler)
	mux.HandleFunc("/api/all", api.GetAllHotListHandler)
	log.Fatal(http.ListenAndServe(":7490", mux))
}

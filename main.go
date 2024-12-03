package main

import (
	"HotSearch/api"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/news/", api.GetHotListHandler)
	log.Fatal(http.ListenAndServe(":7490", mux))
}

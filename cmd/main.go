package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pathespe/kitsune/pkg/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.GetAllApps).Methods(http.MethodGet)
	log.Println("yoa")
	http.ListenAndServe(":3000", router)
}

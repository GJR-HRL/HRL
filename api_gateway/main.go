package main

import (
	"log"
	"net/http"

	"github.com/GJR-HRL/HRL/internal/handlers"
)


func v1() http.Handler{
	v1 := http.NewServeMux()
	v1.HandleFunc("/health", handlers.HealthCheck)
	return v1

}	

func main() {
	router := http.NewServeMux()
	router.Handle("/v1/", http.StripPrefix("/v1",v1()))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}



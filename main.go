package main

import (
	"Go-Tutorial-Api/internal/handlers"
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO Server...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}

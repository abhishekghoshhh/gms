package main

import (
	"net/http"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	capabilitiesController := api.CapabilitiesController()
	router.HandleFunc("/capabilities", capabilitiesController.GetTemplate)
	http.ListenAndServe("0.0.0.0:8081", router)
}

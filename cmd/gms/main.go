package main

import (
	"net/http"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/gorilla/mux"
)

func main() {
	config := config.New()
	println(config.GetString("app.profiles"))
	println(config.Get("age"))
	router := mux.NewRouter()
	capabilitiesController := api.CapabilitiesController()
	router.HandleFunc("/capabilities", capabilitiesController.GetTemplate)
	http.ListenAndServe("0.0.0.0:8081", router)
}

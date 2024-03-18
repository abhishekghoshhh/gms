package main

import (
	"net/http"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/abhishekghoshhh/gms/pkg/lib"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"github.com/gorilla/mux"
)

func main() {
	config := config.New()

	capabilityConfig := *model.NewCapabitiesConfig(
		*model.Entry("scheme", config.FromEnvOrConfig("TOMCAT_CONNECTOR_SCHEME", "scheme")),
		*model.Entry("proxyName", config.FromEnvOrConfig("TOMCAT_CONNECTOR_PROXY_NAME", "proxyName")),
		*model.Entry("proxyPort", config.FromEnvOrConfig("TOMCAT_CONNECTOR_PROXY_PORT", "proxyPort")),
	)

	capabilityBuilder := lib.DefaultCapabilites(capabilityConfig)
	capabilitiesController := api.Capabilities(capabilityBuilder)

	router := mux.NewRouter()
	router.HandleFunc("/capabilities", capabilitiesController.GetTemplate)
	http.ListenAndServe("0.0.0.0:8081", router)
}

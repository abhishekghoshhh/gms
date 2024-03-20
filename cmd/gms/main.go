package main

import (
	"net/http"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/abhishekghoshhh/gms/pkg/client"
	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/abhishekghoshhh/gms/pkg/lib"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"github.com/gorilla/mux"
)

const (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = "8080"
)

func main() {
	config := config.New()
	iamClient := client.New(
		config.FromEnv("IAM_HOST"),
		config.GetString("iam.currentUser"),
	)

	capabilityConfig := *model.NewCapabilitiesConfig(
		*model.Entry("scheme", config.FromEnvOrConfig("TOMCAT_CONNECTOR_SCHEME", "scheme")),
		*model.Entry("proxyName", config.FromEnvOrConfig("TOMCAT_CONNECTOR_PROXY_NAME", "proxyName")),
		*model.Entry("proxyPort", config.FromEnvOrConfig("TOMCAT_CONNECTOR_PROXY_PORT", "proxyPort")),
	)

	passwordGrantflowConfig := model.NewPasswordGrantFlowConfig(
		config.FromEnvOrDefault("PASSWORD_GRANT_FLOW_ACTIVE", "true"),
		config.FromEnv("PASSWORD_GRANT_FLOW_USERNAME"),
		config.FromEnv("PASSWORD_GRANT_FLOW_PASSWORD"),
		config.FromEnv("PASSWORD_GRANT_FLOW_CLIENT_ID"),
		config.FromEnv("PASSWORD_GRANT_FLOW_CLIENT_SECRET"),
	)

	capabilityBuilder := lib.DefaultCapabilities(capabilityConfig)
	capabilitiesController := api.Capabilities(capabilityBuilder)

	gmsService := lib.GmsService(
		passwordGrantflowConfig.IsActive(),
		lib.NewTokenFlow(iamClient),
		lib.NewClientCredentialFlow(),
		lib.NewPasswordGrantFlow(passwordGrantflowConfig),
	)
	gmsController := api.GroupMembership(gmsService)

	router := mux.NewRouter()
	router.HandleFunc("/capabilities", capabilitiesController.GetTemplate)
	router.HandleFunc("/gms/search", gmsController.GetGroups)
	http.ListenAndServe(SERVER_HOST+":"+SERVER_PORT, router)
}

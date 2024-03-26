package main

import (
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/client"
	"github.com/abhishekghoshhh/gms/pkg/httpclient"

	"github.com/abhishekghoshhh/gms/internal/api"
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
	httpClinet := httpclient.NewClient()
	iamClient := client.New(
		httpClinet,
		config.FromEnv("IAM_HOST"),
		config.GetString("iam.currentUser"),
	)

	capabilityConfig := *model.NewCapabilitiesConfig(
		*model.NewEntry("scheme", config.FromEnvOrConfig("TOMCAT_CONNECTOR_SCHEME", "scheme")),
		*model.NewEntry("proxyName", config.FromEnvOrConfig("TOMCAT_CONNECTOR_PROXY_NAME", "proxyName")),
		*model.NewEntry("proxyPort", config.FromEnvOrConfig("TOMCAT_CONNECTOR_PROXY_PORT", "proxyPort")),
	)

	capabilityBuilder := lib.CapabilitiesBuilder(capabilityConfig)
	capabilitiesApi := api.Capabilities(capabilityBuilder)

	authTokenFlow := lib.NewAuthTokenFlow(iamClient)
	clientCertFlow := lib.NewClientCertFlow(
		config.FromEnvOrDefault("PASSWORD_GRANT_FLOW_ACTIVE", "true"),
	)
	gmsService := lib.GmsService(authTokenFlow, clientCertFlow)
	gmsApi := api.GroupMembershipCheck(gmsService)

	router := mux.NewRouter()
	router.HandleFunc("/capabilities", capabilitiesApi.GetTemplate)
	router.HandleFunc("/gms/search", gmsApi.GetGroups)
	http.ListenAndServe(SERVER_HOST+":"+SERVER_PORT, router)
}

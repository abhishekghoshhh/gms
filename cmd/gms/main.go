package main

import (
	"net/http"
	"os"
	"time"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/abhishekghoshhh/gms/pkg/client"
	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/abhishekghoshhh/gms/pkg/httpclient"
	"github.com/abhishekghoshhh/gms/pkg/lib"
	"github.com/abhishekghoshhh/gms/pkg/logger"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"github.com/gorilla/mux"
)

const (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = "8080"
)

func main() {
	config := config.New()
	timeout := 2 * time.Second

	httpClient := httpclient.NewClient(timeout)

	iamClient, err := client.New(
		httpClient,
		os.Getenv("IAM_HOST"),
		config.GetString("iam.apis.currentUser"),
		config.GetString("iam.apis.fetchToken"),
		config.GetString("iam.apis.findAccountByCertSubject"),
		config.GetString("iam.apis.getUserCount"),
		config.GetString("iam.apis.fetchUsersInBatch"),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	capabilityConfig := *model.NewCapabilitiesConfig(
		*model.NewEntry("scheme", config.FromEnvOrConfig("TOMCAT_CONNECTOR_SCHEME", "scheme")),
		*model.NewEntry("proxyName", config.FromEnvOrConfig("TOMCAT_CONNECTOR_PROXY_NAME", "proxyName")),
		*model.NewEntry("proxyPort", config.FromEnvOrConfig("TOMCAT_CONNECTOR_PROXY_PORT", "proxyPort")),
	)

	capabilityBuilder := lib.CapabilitiesBuilder(capabilityConfig)
	capabilitiesApi := api.Capabilities(capabilityBuilder)

	passwordGrantConfig := model.NewPasswordGrantFlowConfig(
		os.Getenv("PASSWORD_GRANT_FLOW_USERNAME"),
		os.Getenv("PASSWORD_GRANT_FLOW_PASSWORD"),
		os.Getenv("PASSWORD_GRANT_FLOW_CLIENT_ID"),
		os.Getenv("PASSWORD_GRANT_FLOW_CLIENT_SECRET"),
	)

	clientCredentialConfig := model.NewClientCredentialConfig(
		config.FromEnvOrConfig("SCIM_LIST_USER_BATCH_SIZE", "scim-read-client.batchSize"),
		config.FromEnvOrConfig("SCIM_LIST_USER_MAX_BATCH_COUNT", "scim-read-client.maxBatchCount"),
		os.Getenv("SCIM_READ_CLIENT_ID"),
		os.Getenv("SCIM_READ_CLIENT_SECRET"),
	)

	authTokenFlow := lib.NewAuthTokenFlow(iamClient)
	clientCertFlow, err := lib.NewClientCertFlow(
		iamClient,
		config.FromEnvOrConfig("PASSWORD_GRANT_FLOW_ACTIVE", "password-grant-flow.isEnabled"),
		passwordGrantConfig,
		clientCredentialConfig,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	gmsService := lib.GmsService(authTokenFlow, clientCertFlow)
	gmsApi := api.GroupMembershipCheck(gmsService)

	router := mux.NewRouter()
	router.HandleFunc("/capabilities", capabilitiesApi.GetTemplate)
	router.HandleFunc("/gms/search", gmsApi.GetGroups)

	http.ListenAndServe(SERVER_HOST+":"+SERVER_PORT, router)
}

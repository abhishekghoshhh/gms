package main

import (
	"fmt"
	"github.com/abhishekghoshhh/gms/pkg/handlers"
	"github.com/abhishekghoshhh/gms/pkg/iamclient"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"os"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/abhishekghoshhh/gms/pkg/client"
	"github.com/abhishekghoshhh/gms/pkg/config"
	httpclient "github.com/abhishekghoshhh/gms/pkg/http"
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
	c := config.New()
	fmt.Println(c.GetString("iam.clientCredentialToken.clientSecret"))
	fmt.Println(c.GetString("iam.clientCredentialToken.clientId"))

	var configs iamclient.IamConfigs
	err := mapstructure.Decode(c.Get("iam"), &configs.Config)

	if err != nil {
		logger.Fatal("error in loading config")
	}

	client1 := iamclient.NewIamClient(os.Getenv("IAM_HOST"), &configs)
	handler := handlers.NewHandler(*client1)

	router := mux.NewRouter()
	router.HandleFunc("/gms/search", handler.GetGroups)
}

func mainCopy() {
	config := config.New()

	httpClient := httpclient.NewClient()

	iamClient, err := client.New(
		httpClient,
		os.Getenv("IAM_HOST"),
		config.GetString("iam.apis.currentUser"),
		config.GetString("iam.apis.fetchToken"),
		config.GetString("iam.apis.findAccountByCertSubject"),
		config.GetString("iam.apis.getUserCount"),
		config.GetString("iam.apis.fetchUsersInBatch"),
		config.GetString("iam.apis.fetchUserById"),
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

	logger.Info("GMS application is running on " + SERVER_HOST + ":" + SERVER_PORT)
	http.ListenAndServe(SERVER_HOST+":"+SERVER_PORT, router)
}

package main

import (
	"os"
	"path/filepath"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/labstack/echo"

	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/abhishekghoshhh/gms/pkg/http"
	"github.com/abhishekghoshhh/gms/pkg/iam"
	"github.com/abhishekghoshhh/gms/pkg/logger"
)

func main() {
	logger.Debug("I am here")
	c := config.New()
	e := echo.New()

	serverPort := c.GetString("server.port")

	iamConfig := make(map[string]*iam.IamConfig)
	c.Decode("iam", &iamConfig)

	iamClient := iam.New(
		os.Getenv("IAM_HOST"),
		iamConfig,
		http.NewClient(),
	)

	workingDir, _ := os.Getwd()
	capabilitiesPath := filepath.Join(
		workingDir,
		c.GetString("server.capabilities.path"),
	)
	capabilitiesConfig := make(map[string]string)
	c.Decode("server.capabilities.config", &capabilitiesConfig)

	groupsHandler := api.NewGetGroupsHandler(iamClient)
	capabilitiesHandler := api.CapabilitiesHandler(capabilitiesConfig, capabilitiesPath)

	e.GET("/gms/search", groupsHandler.GetGroups)
	e.GET("/gms/capabilities", capabilitiesHandler.GetTemplate)

	e.Logger.Fatal(e.Start(":" + serverPort))
}

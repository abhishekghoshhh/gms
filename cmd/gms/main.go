package main

import (
	"os"
	"path/filepath"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/abhishekghoshhh/gms/pkg/http"
	"github.com/abhishekghoshhh/gms/pkg/iam"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// creating an instance of the config
	cfg := config.New()

	// building dependencies for GMS api
	var iamConfig config.IamConfig
	cfg.Decode("iam", &iamConfig)
	iamClient := iam.New(
		iamConfig,
		http.NewClient(),
	)
	groupsHandler := api.NewGetGroupsHandler(iamClient)

	// building dependencies for capabilities api
	workingDir, _ := os.Getwd()
	capabilitiesPath := filepath.Join(
		workingDir,
		cfg.GetString("server.capabilities.path"),
	)
	capabilitiesConfig := make(map[string]string)
	cfg.Decode("server.capabilities.config", &capabilitiesConfig)
	capabilitiesHandler := api.CapabilitiesHandler(capabilitiesConfig, capabilitiesPath)

	// hosting dynamic apis
	e.GET("/gms/search", groupsHandler.GetGroups)
	e.GET("/gms/capabilities", capabilitiesHandler.GetTemplate)

	// hosting static files
	e.File("/", cfg.GetString("server.web.index"))
	e.File("/swagger.yaml", cfg.GetString("server.swagger"))
	e.Static("/swagger-ui/*", cfg.GetString("server.web.static"))

	// spinning up the server
	e.Logger.SetLevel(log.INFO)
	e.Logger.Fatal(e.Start(":" + cfg.GetString("server.port")))
}

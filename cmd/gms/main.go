package main

import (
	"os"

	"github.com/abhishekghoshhh/gms/internal/api"
	"github.com/labstack/echo"

	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/abhishekghoshhh/gms/pkg/iam"
)

const (
	SERVER_PORT = "8080"
)

func main() {
	c := config.New()

	iamConfig := make(map[string]*iam.IamConfig)
	c.Decode("iam", &iamConfig)

	iamClient := iam.New(
		os.Getenv("IAM_HOST"),
		iamConfig,
	)

	handler := api.NewHandler(iamClient)

	e := echo.New()

	e.GET("/gms/search", handler.GetGroups)
	e.Logger.Fatal(e.Start(":" + SERVER_PORT))
}

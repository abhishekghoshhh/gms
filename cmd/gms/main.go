package main

import (
	"os"

	"github.com/abhishekghoshhh/gms/internal/api"

	"github.com/abhishekghoshhh/gms/pkg/config"
	"github.com/abhishekghoshhh/gms/pkg/iam"
	"github.com/gorilla/mux"
)

const (
	SERVER_HOST = "0.0.0.0"
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

	router := mux.NewRouter()
	router.HandleFunc("/gms/search", handler.GetGroups)
}

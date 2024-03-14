package api

import (
	"fmt"
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/lib"
)

type capabilities struct {
}

func CapabilitiesController() *capabilities {
	return &capabilities{}
}

func (*capabilities) GetTemplate(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "application/xml")
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprintf(responseWriter, lib.CapabilitiesTemplate())
}

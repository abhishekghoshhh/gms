package api

import (
	"fmt"
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/lib"
)

type CapabilitiesApi struct {
	capabilityBuilder lib.CapabilityBuilder
}

func Capabilities(capabilityBuilder lib.CapabilityBuilder) *CapabilitiesApi {
	return &CapabilitiesApi{
		capabilityBuilder,
	}
}

func (capabilitesApi *CapabilitiesApi) GetTemplate(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Add("Content-Type", "application/xml")
	responseWriter.WriteHeader(http.StatusOK)
	fmt.Fprint(responseWriter, capabilitesApi.capabilityBuilder.Capabilities())
}

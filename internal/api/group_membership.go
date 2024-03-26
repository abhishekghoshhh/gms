package api

import (
	"fmt"
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/lib"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

type GroupMembershipApi struct {
	gmsFlow lib.GmsFlow
}

func (gmsApi *GroupMembershipApi) GetGroups(responseWriter http.ResponseWriter, request *http.Request) {
	gms := model.GMS(
		request.Header.Get("Authorization"),
		request.Header.Get("X-SSL-Client-S-Dn"),
		request.Header.Get("X-SSL-Client-Cert"),
		request.URL.Query()["group"],
	)
	if resp, err := gmsApi.gmsFlow.GetGroups(gms); err != nil {
		responseWriter.Header().Add("Content-Type", "text/plain")
		responseWriter.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(responseWriter, err)
	} else {
		responseWriter.Header().Add("Content-Type", "text/plain")
		responseWriter.WriteHeader(http.StatusOK)
		fmt.Fprint(responseWriter, resp)
	}
}

func GroupMembershipCheck(gmsFlow lib.GmsFlow) *GroupMembershipApi {
	return &GroupMembershipApi{
		gmsFlow,
	}
}

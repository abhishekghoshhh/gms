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
	responseWriter.Header().Add("Content-Type", "text/plain")
	responseWriter.WriteHeader(http.StatusOK)
	gms := model.GMS()
	resp, _ := gmsApi.gmsFlow.GetGroups(gms)
	fmt.Fprint(responseWriter, resp)
}

func GroupMembership(gmsFlow lib.GmsFlow) *GroupMembershipApi {
	return &GroupMembershipApi{
		gmsFlow,
	}
}

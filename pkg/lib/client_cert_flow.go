package lib

import (
	"os"

	"github.com/abhishekghoshhh/gms/pkg/model"
	"github.com/abhishekghoshhh/gms/pkg/util"
)

type ClientCertFlow struct {
	isPasswordGrantFlowActive bool
}

func (flow *ClientCertFlow) GetGroups(data *model.GmsModel) (string, error) {
	if flow.isPasswordGrantFlowActive {
		return flow.fromPasswordGrant(data)
	} else {
		return flow.fromClientCredential(data)
	}
}

func (flow *ClientCertFlow) fromClientCredential(data *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}

func (flow *ClientCertFlow) fromPasswordGrant(data *model.GmsModel) (string, error) {
	_ = os.Getenv("PASSWORD_GRANT_FLOW_USERNAME")
	_ = os.Getenv("PASSWORD_GRANT_FLOW_PASSWORD")
	_ = os.Getenv("PASSWORD_GRANT_FLOW_CLIENT_ID")
	_ = os.Getenv("PASSWORD_GRANT_FLOW_CLIENT_SECRET")
	return "group1\ngroup2", nil
}

func NewClientCertFlow(isPasswordGrantFlowActive string) *ClientCertFlow {
	return &ClientCertFlow{
		isPasswordGrantFlowActive: util.Bool(isPasswordGrantFlowActive),
	}
}

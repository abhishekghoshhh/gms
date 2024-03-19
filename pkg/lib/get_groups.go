package lib

import "github.com/abhishekghoshhh/gms/pkg/model"

type GmsFlow interface {
	GetGroups(flowType model.Flow) (string, error)
}

type GmsFlowService struct {
	tokenFlow          GmsFlow
	clientCrendialFlow GmsFlow
	passwordGrantFlow  GmsFlow
	defaultGmsFlow     GmsFlow
}

func (gmsService *GmsFlowService) GetGroups(flowType model.Flow) (string, error) {
	return gmsService.getFlow(flowType).GetGroups(flowType)
}

func (gmsService *GmsFlowService) getFlow(flowType model.Flow) GmsFlow {
	switch flowType.GetFlowType() {
	case model.TOKEN_FLOW:
		return gmsService.tokenFlow
	case model.CLIENT_CREDENTIAL_FLOW:
		return gmsService.clientCrendialFlow
	case model.PASSWORD_GRANT_FLOW:
		return gmsService.passwordGrantFlow
	default:
		return gmsService.defaultGmsFlow
	}
}

func GmsService(tokenFlow, clientCredentialFlow, passwordGrantFlow, defaultGmsFlow GmsFlow) *GmsFlowService {
	return &GmsFlowService{
		tokenFlow,
		clientCredentialFlow,
		passwordGrantFlow,
		defaultGmsFlow,
	}
}

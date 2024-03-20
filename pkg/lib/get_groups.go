package lib

import (
	"errors"

	"github.com/abhishekghoshhh/gms/pkg/model"
)

type GmsFlow interface {
	GetGroups(gmsModel *model.GmsModel) (string, error)
}

type GmsFlowService struct {
	isPasswordGrantFlowEnabled bool
	tokenFlow            GmsFlow
	clientCredentialFlow GmsFlow
	passwordGrantFlow    GmsFlow
}

func (gmsService *GmsFlowService) GetGroups(gmsModel *model.GmsModel) (string, error) {
	if gmsFlow, err := gmsService.getFlow(gmsModel); err != nil {
		return "", err
	} else {
		return gmsFlow.GetGroups(gmsModel)
	}
}

func (gmsService *GmsFlowService) getFlow(gmsModel *model.GmsModel) (GmsFlow, error) {
	if gmsModel.HasToken() {
		return gmsService.tokenFlow, nil
	}
	if gmsModel.HasCert() {
		if gmsService.isPasswordGrantFlowEnabled {
			return gmsService.passwordGrantFlow, nil
		} else {
			return gmsService.clientCredentialFlow, nil
		}
	}
	return nil, errors.New("token and cert both not present")
}

func GmsService(isPasswordGrantFlowEnabled bool, tokenFlow, clientCredentialFlow, passwordGrantFlow GmsFlow) *GmsFlowService {
	return &GmsFlowService{
		isPasswordGrantFlowEnabled,
		tokenFlow,
		clientCredentialFlow,
		passwordGrantFlow,
	}
}

package lib

import (
	"errors"

	"github.com/abhishekghoshhh/gms/pkg/client"
	"github.com/abhishekghoshhh/gms/pkg/logger"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"github.com/abhishekghoshhh/gms/pkg/util"
)

type ClientCertFlow struct {
	iamClient                 *client.IamClient
	isPasswordGrantFlowActive bool
	passwordGrantConfig       *model.PasswordGrantConfig
	clientCredentialConfig    *model.ClientCredentialConfig
}

func (flow *ClientCertFlow) GetGroups(data *model.GmsModel) (string, error) {
	if flow.isPasswordGrantFlowActive {
		return flow.fromPasswordGrant(data)
	} else {
		return flow.fromClientCredential(data)
	}
}

func (flow *ClientCertFlow) fromClientCredential(data *model.GmsModel) (string, error) {
	logger.Info("Client credential flow for GMS group search is executing")

	return "group1\ngroup2", nil
}

func (flow *ClientCertFlow) fromPasswordGrant(data *model.GmsModel) (string, error) {
	logger.Info("Password grant flow for GMS group search is executing")

	return "group1\ngroup2", nil
}

func NewClientCertFlow(iamClient *client.IamClient, isPasswordGrantFlowActive string,
	passwordGrantConfig *model.PasswordGrantConfig, clientCredentialConfig *model.ClientCredentialConfig) (*ClientCertFlow, error) {

	passwordGrantFlowEnabled := util.Bool(isPasswordGrantFlowActive)
	if passwordGrantFlowEnabled && !passwordGrantConfig.IsValid() {
		return nil, errors.New("password grant flow is active but password grant config is not valid")
	} else if !passwordGrantFlowEnabled && !clientCredentialConfig.IsValid() {
		return nil, errors.New("password grant flow is inactive but client credential config is not valid")
	}

	return &ClientCertFlow{
		iamClient:                 iamClient,
		isPasswordGrantFlowActive: passwordGrantFlowEnabled,
		passwordGrantConfig:       passwordGrantConfig,
		clientCredentialConfig:    clientCredentialConfig,
	}, nil
}

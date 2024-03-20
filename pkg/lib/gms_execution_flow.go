package lib

import "github.com/abhishekghoshhh/gms/pkg/model"

type TokenFlow struct {
}

func NewTokenFlow() *TokenFlow {
	return &TokenFlow{}
}

func (flow *TokenFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}

type ClientCredentialFlow struct {
}

func NewClientCredentialFlow() *ClientCredentialFlow {
	return &ClientCredentialFlow{}
}

func (flow *ClientCredentialFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}

type PasswordGrantFlow struct {
	cfg *model.PasswordGrantFlowConfig
}

func NewPasswordGrantFlow(cfg *model.PasswordGrantFlowConfig) *PasswordGrantFlow {
	return &PasswordGrantFlow{
		cfg,
	}
}

func (flow *PasswordGrantFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}

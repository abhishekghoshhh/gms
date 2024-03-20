package lib

import "github.com/abhishekghoshhh/gms/pkg/model"

type TokenFlow struct {
}

func (flow *TokenFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}

type ClientCredentialFlow struct {
}

func (flow *ClientCredentialFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}

type PasswordGrantFlow struct {
}

func (flow *PasswordGrantFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}

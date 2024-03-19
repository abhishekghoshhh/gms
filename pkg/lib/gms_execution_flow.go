package lib

import "github.com/abhishekghoshhh/gms/pkg/model"

type TokenFlow struct {
}

func (flow *TokenFlow) GetGroups(flowType model.Flow) (string, error) {
	return "group1\ngroup2", nil
}

type ClientCredentialFlow struct {
}

func (flow *ClientCredentialFlow) GetGroups(flowType model.Flow) (string, error) {
	return "group1\ngroup2", nil
}

type PasswordGrantFlow struct {
}

func (flow *PasswordGrantFlow) GetGroups(flowType model.Flow) (string, error) {
	return "group1\ngroup2", nil
}

type DefaultGmsFlow struct {
}

func (flow *DefaultGmsFlow) GetGroups(flowType model.Flow) (string, error) {
	return "default-group1\ndefault-group2", nil
}

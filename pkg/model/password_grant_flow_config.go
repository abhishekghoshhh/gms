package model

import (
	"strconv"

	"github.com/abhishekghoshhh/gms/pkg/logger"
)

type PasswordGrantFlowConfig struct {
	isActive     bool
	userName     string
	password     string
	clientId     string
	clientSecret string
}

func NewPasswordGrantFlowConfig(isActive, userName, password, clientId, clientSecret string) *PasswordGrantFlowConfig {
	logger.Debug("isPasswordGrantFlowActive : " + isActive)
	return &PasswordGrantFlowConfig{
		isActive:     Bool(isActive),
		userName:     userName,
		password:     password,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func Bool(val string) bool {
	if b, err := strconv.ParseBool(val); err != nil {
		return false
	} else {
		return b
	}
}

func (cfg *PasswordGrantFlowConfig) IsActive() bool {
	return cfg.isActive
}

package model

import "github.com/abhishekghoshhh/gms/pkg/util"

type ClientCredentialConfig struct {
	batchSize     int
	maxBatchCount int
	clientId      string
	clientSecret  string
}

func (config *ClientCredentialConfig) IsValid() bool {
	return config.batchSize != 0 && config.maxBatchCount != 0 && config.clientId != "" && config.clientSecret != ""
}
func (config *ClientCredentialConfig) BatchSize() int {
	return config.batchSize
}
func (config *ClientCredentialConfig) MaxBatchCount() int {
	return config.maxBatchCount
}
func (config *ClientCredentialConfig) ClientId() string {
	return config.clientId
}
func (config *ClientCredentialConfig) ClientSecret() string {
	return config.clientSecret
}

func NewClientCredentialConfig(batchSize, maxBatchCount, clientId, clientSecret string) *ClientCredentialConfig {
	return &ClientCredentialConfig{
		batchSize:     util.Int(batchSize),
		maxBatchCount: util.Int(maxBatchCount),
		clientId:      clientId,
		clientSecret:  clientSecret,
	}
}

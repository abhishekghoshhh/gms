package model

type PasswordGrantConfig struct {
	userName     string
	password     string
	clientId     string
	clientSecret string
}

func (config *PasswordGrantConfig) IsValid() bool {
	return (config.userName != "" && config.password != "" && config.clientId != "" && config.clientSecret != "")
}

func (config *PasswordGrantConfig) UserName() string {
	return config.userName
}

func (config *PasswordGrantConfig) Password() string {
	return config.password
}

func (config *PasswordGrantConfig) ClientId() string {
	return config.clientId
}

func (config *PasswordGrantConfig) ClientSecret() string {
	return config.clientSecret
}

func NewPasswordGrantFlowConfig(userName, password, clientId, clientSecret string) *PasswordGrantConfig {
	return &PasswordGrantConfig{
		userName:     userName,
		password:     password,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

package config

type ApiConfig struct {
	Path         string
	Timeout      int
	ClientId     string
	ClientSecret string
}

type IamApiConfig struct {
	UserInfo              ApiConfig
	ClientCredentialToken ApiConfig
	FetchUserById         ApiConfig
}

type IamConfig struct {
	Host string
	Apis IamApiConfig
}

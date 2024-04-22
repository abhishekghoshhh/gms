package iamclient

import (
	"errors"
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/config"
	httpclient "github.com/abhishekghoshhh/gms/pkg/http"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

const (
	MediaTypeAll          = "*/*"
	ApplicationUrlEncoded = "x-www-form-urlencoded"
)

type IamConfigs struct {
	Config map[string]*IamConfig `mapstructure:"iam"`
}
type IamConfig struct {
	*config.ApiConfig
	ClientId     string
	ClientSecret string
}

type IamClient struct {
	Host   string
	Config map[string]*IamConfig
	client *httpclient.CustomClient
}

func NewIamClient(host string, iamConfigs *IamConfigs) *IamClient {
	return &IamClient{
		Host:   host,
		Config: iamConfigs.Config,
	}
}

func (iamClient *IamClient) FetchUserById(token, userId string) (*model.IamProfileResponse, error) {
	apiConfig, ok := iamClient.Config["fetchUserById"]
	if !ok {
		return nil, errors.New("invalid Api Config")
	}

	headers := map[string]string{
		"Authorization": token,
		"Accept":        MediaTypeAll,
	}
	queryParams := map[string]string{
		"userId": userId,
	}

	req := httpclient.
		Request(
			iamClient.Host,
			apiConfig.Path,
			http.MethodGet,
		).
		Headers(headers).
		QueryParams(queryParams).
		Timeout(apiConfig.Timeout)

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.IamProfileResponse{})
	}
}

func (iamClient *IamClient) FetchClientCredentialToken() (*model.ClientTokenResponse, error) {
	apiConfig := iamClient.Config["clientCredentialToken"]

	request := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     apiConfig.ClientId,
		"client_secret": apiConfig.ClientSecret,
	}
	return iamClient.getToken(request)
}
func (iamClient *IamClient) getToken(requestBody map[string]string) (*model.ClientTokenResponse, error) {
	apiConfig := iamClient.Config["clientCredentialToken"]

	headers := map[string]string{
		"Accept":       MediaTypeAll,
		"Content-Type": ApplicationUrlEncoded,
	}

	req := httpclient.
		Request(
			iamClient.Host,
			apiConfig.Path,
			http.MethodPost,
		).
		Headers(headers).
		Body(requestBody).
		Timeout(apiConfig.Timeout)

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.ClientTokenResponse{})
	}
}
func (iamClient *IamClient) FetchUserInfo(token string) (*model.UserInfo, error) {
	apiConfig := iamClient.Config["userinfo"]

	headers := map[string]string{
		"Accept":        MediaTypeAll,
		"Content-Type":  ApplicationUrlEncoded,
		"Authorization": token,
	}

	req := httpclient.
		Request(
			iamClient.Host,
			apiConfig.Path,
			http.MethodPost,
		).
		Headers(headers).
		Timeout(apiConfig.Timeout)

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.UserInfo{})
	}
}

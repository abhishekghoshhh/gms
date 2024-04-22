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

func (iamClient *IamClient) FetchAccessTokenForClientCredentialFlow(clientId, clientSecret string) (*model.ClientTokenResponse, error) {
	request := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     clientId,
		"client_secret": clientSecret,
	}
	return iamClient.getBearerToken(request)
}
func (iamClient *IamClient) getBearerToken(requestBody map[string]string) (*model.ClientTokenResponse, error) {
	headers := map[string]string{
		"Accept":       MediaTypeAll,
		"Content-Type": ApplicationUrlEncoded,
	}

	req := httpclient.
		Request(
			iamClient.iamHost,
			iamClient.tokenApi,
			http.MethodPost,
		).
		Headers(headers).
		Body(requestBody).
		Timeout(2) //timeout needs to be changed

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.ClientTokenResponse{})
	}
}

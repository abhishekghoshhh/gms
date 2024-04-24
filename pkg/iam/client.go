package iam

import (
	"errors"
	"fmt"
	"net/http"

	httpclient "github.com/abhishekghoshhh/gms/pkg/http"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

const (
	MediaTypeAll          = "*/*"
	ApplicationUrlEncoded = "x-www-form-urlencoded"
)

type IamConfig struct {
	Path         string
	Timeout      int
	ClientId     string
	ClientSecret string
}

type IamClient struct {
	Host   string
	Config map[string]*IamConfig
	Client *httpclient.CustomClient
}

func New(host string, iamConfig map[string]*IamConfig, client *httpclient.CustomClient) *IamClient {
	return &IamClient{
		Host:   host,
		Config: iamConfig,
		Client: client,
	}
}

func (iamClient *IamClient) FetchUserById(token, userId string) (*model.IamProfileResponse, error) {
	apiConfig, ok := iamClient.Config["fetchuserbyid"]
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

	if resp, err := iamClient.Client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.IamProfileResponse{})
	}
}

func (iamClient *IamClient) FetchClientCredentialToken() (*model.Token, error) {
	apiConfig := iamClient.Config["clientcredentialtoken"]

	request := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     apiConfig.ClientId,
		"client_secret": apiConfig.ClientSecret,
	}
	return iamClient.getToken(request)
}
func (iamClient *IamClient) getToken(requestBody map[string]string) (*model.Token, error) {
	apiConfig := iamClient.Config["clientcredentialtoken"]

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
		Timeout(apiConfig.Timeout).
		Log()

	if resp, err := iamClient.Client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.Token{})
	}
}
func (iamClient *IamClient) FetchUserInfo(token string) (*model.UserInfo, error) {
	fmt.Println("iamClient.Config", iamClient.Config)
	apiConfig := iamClient.Config["userinfo"]
	fmt.Println("apiConfig", apiConfig)

	headers := map[string]string{
		"Accept":        MediaTypeAll,
		"Content-Type":  ApplicationUrlEncoded,
		"Authorization": token,
	}

	req := httpclient.
		Request(
			iamClient.Host,
			apiConfig.Path,
			http.MethodGet,
		).
		Headers(headers).
		Timeout(apiConfig.Timeout)

	if resp, err := iamClient.Client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.UserInfo{})
	}
}

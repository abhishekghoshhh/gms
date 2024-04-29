package iam

import (
	"errors"
	"net/http"
	"net/url"

	httpclient "github.com/abhishekghoshhh/gms/pkg/http"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

const (
	MediaTypeAll          = "*/*"
	ApplicationUrlEncoded = "application/x-www-form-urlencoded"
)

type IamConfig struct {
	Path         string
	Timeout      int
	ClientId     string
	ClientSecret string
}
type Client interface {
	FetchUserById(token, userId string) (*model.IamProfileResponse, error)
	FetchClientCredentialToken() (*model.Token, error)
	FetchUserInfo(token string) (*model.UserInfo, error)
}

type IamClient struct {
	Host   string
	Config map[string]IamConfig
	Client *httpclient.CustomClient
}

func New(host string, iamConfig map[string]IamConfig, client *httpclient.CustomClient) *IamClient {
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
		"Authorization": "Bearer " + token,
		"Accept":        MediaTypeAll,
	}
	pathVariables := map[string]string{
		"userId": userId,
	}

	req := httpclient.
		Request(
			iamClient.Host,
			apiConfig.Path,
			http.MethodGet,
		).
		Headers(headers).
		PathVariables(pathVariables).
		Timeout(apiConfig.Timeout)

	if resp, err := iamClient.Client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.IamProfileResponse{})
	}
}

func (iamClient *IamClient) FetchClientCredentialToken() (*model.Token, error) {
	apiConfig := iamClient.Config["clientcredentialtoken"]

	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", apiConfig.ClientId)
	formData.Set("client_secret", apiConfig.ClientSecret)

	return iamClient.getToken(formData.Encode())
}

func (iamClient *IamClient) getToken(requestBody any) (*model.Token, error) {
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
		UrlEncodedData().
		Log()

	if resp, err := iamClient.Client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.Token{})
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

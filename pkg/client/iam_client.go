package client

import (
	"errors"

	"github.com/abhishekghoshhh/gms/pkg/httpclient"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

type IamClient struct {
	iamHost        string
	scimProfileApi string
	client         *httpclient.Client
}

func New(client *httpclient.Client, iamHost, scimProfileApi string) (*IamClient, error) {
	if iamHost == "" {
		return nil, errors.New("iam host is null, please set the iam host")
	}
	return &IamClient{
		iamHost:        iamHost,
		scimProfileApi: scimProfileApi,
		client:         client,
	}, nil
}

func (iamClient *IamClient) FetchUser(token string) (*model.IamProfileResponse, error) {
	headers := make(map[string]string)
	headers["Authorization"] = token
	headers["Accept"] = "*/*"

	request, err := iamClient.client.Create("GET", iamClient.iamHost, iamClient.scimProfileApi, headers)
	if err != nil {
		return nil, err
	}

	response, err := iamClient.client.Send(request)
	if err != nil {
		return nil, err
	}
	return httpclient.Parse(response, &model.IamProfileResponse{})
}

func (iamClient *IamClient) FetchAccessTokenForPasswordGrantFlow(username, password, clientId, clientSecret string) (string, error) {
	headers := make(map[string]string)
	headers["grant_type"] = "password"
	headers["username"] = username
	headers["password"] = password
	headers["client_id"] = clientId
	headers["client_secret"] = clientSecret

	return iamClient.getBearerToken(headers)
}

func (iamClient *IamClient) getBearerToken(headers map[string]string) (string, error) {
	headers["Accept"] = "*/*"
	headers["Content-Type"] = "x-www-form-urlencoded"

	request, err := iamClient.client.Create("POST", iamClient.iamHost, "/token", headers)
	if err != nil {
		return "", err
	}

	resp, err := iamClient.client.Send(request)
	if err != nil {
		return "", err
	}

	parsedResponse, err := httpclient.Parse(resp, &model.ClientTokenResponse{})
	if err != nil {
		return "", err
	}

	if iamClient.isBearerToken(parsedResponse) {
		return parsedResponse.AccessToken, nil
	}

	return "", errors.New("unable to retrieve Access Token for flow")
}

func (iamClient *IamClient) isBearerToken(response *model.ClientTokenResponse) bool {
	return response != nil && response.AccessToken != ""
}

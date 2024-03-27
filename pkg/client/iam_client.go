package client

import (
	"github.com/abhishekghoshhh/gms/pkg/httpclient"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

type IamClient struct {
	iamHost        string
	scimProfileApi string
	client         *httpclient.Client
}

func New(client *httpclient.Client, iamHost, scimProfileApi string) *IamClient {
	return &IamClient{
		iamHost:        iamHost,
		scimProfileApi: scimProfileApi,
		client:         client,
	}
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

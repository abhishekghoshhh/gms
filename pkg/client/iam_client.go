package client

import (
	"encoding/json"
	"errors"
	"github.com/abhishekghoshhh/gms/pkg/httpclient"
	"github.com/abhishekghoshhh/gms/pkg/logger"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"go.uber.org/zap"
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

	var iamProfileResponse model.IamProfileResponse
	if err := json.Unmarshal(response, &iamProfileResponse); err != nil {
		logger.Error("error is " + err.Error())
		return nil, errors.New("invalid profileResponse")
	}
	logger.Info("profile response is ", zap.Any("resp", iamProfileResponse))
	return &iamProfileResponse, nil
}

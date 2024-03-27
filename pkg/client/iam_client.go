package client

import (
	"errors"
	"strconv"

	"github.com/abhishekghoshhh/gms/pkg/httpclient"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

const (
	GET  = "GET"
	POST = "POST"
)

type IamClient struct {
	iamHost                     string
	scimProfileApi              string
	tokenApi                    string
	findAccountByCertSubjectApi string
	getUserCountApi             string
	fetchUsersInBatchApi        string
	client                      *httpclient.Client
}

func New(client *httpclient.Client, iamHost,
	scimProfileApi, tokenApi, findAccountByCertSubjectApi, getUserCountApi, fetchUsersInBatchApi string) (*IamClient, error) {

	if iamHost == "" {
		return nil, errors.New("iam host is null, please set the iam host")
	}
	return &IamClient{
		iamHost,
		scimProfileApi,
		tokenApi,
		findAccountByCertSubjectApi,
		getUserCountApi,
		fetchUsersInBatchApi,
		client,
	}, nil
}

func (iamClient *IamClient) FetchUser(token string) (*model.IamProfileResponse, error) {
	headers := map[string]string{
		"Authorization": token,
		"Accept":        "*/*",
	}

	request, err := iamClient.client.Create(GET, iamClient.iamHost, iamClient.scimProfileApi, headers)
	if err != nil {
		return nil, err
	}

	response, err := iamClient.client.Send(request)
	if err != nil {
		return nil, err
	}
	return httpclient.Parse(response, &model.IamProfileResponse{})
}

func (iamClient *IamClient) FetchUserCount(token string) (*model.IamProfileListResponse, error) {
	headers := map[string]string{
		"Authorization": token,
		"Accept":        "*/*",
	}
	request, err := iamClient.client.Create(GET, iamClient.iamHost, iamClient.getUserCountApi, headers)
	if err != nil {
		return nil, err
	}

	response, err := iamClient.client.Send(request)
	if err != nil {
		return nil, err
	}
	return httpclient.Parse(response, &model.IamProfileListResponse{})
}

func (iamClient *IamClient) FetchUsersInBatch(token string, startingIndex, count int) (*model.IamProfileListResponse, error) {
	headers := map[string]string{
		"Authorization": token,
		"Accept":        "*/*",
	}

	queryParams := map[string]string{
		"startIndex": strconv.Itoa(startingIndex),
		"count":      strconv.Itoa(count),
	}
	request, err := iamClient.client.CreateWithParams(GET, iamClient.iamHost, iamClient.fetchUsersInBatchApi, headers, queryParams, nil)
	if err != nil {
		return nil, err
	}

	response, err := iamClient.client.Send(request)
	if err != nil {
		return nil, err
	}
	return httpclient.Parse(response, &model.IamProfileListResponse{})
}

func (iamClient *IamClient) FetchAccessTokenForClientCredentialFlow(clientId, clientSecret string) (*model.ClientTokenResponse, error) {
	request := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     clientId,
		"client_secret": clientSecret,
	}
	return iamClient.getBearerToken(request)
}

func (iamClient *IamClient) FetchAccessTokenForPasswordGrantFlow(username, password, clientId, clientSecret string) (*model.ClientTokenResponse, error) {
	request := map[string]string{
		"grant_type":    "password",
		"username":      username,
		"password":      password,
		"client_id":     clientId,
		"client_secret": clientSecret,
	}
	return iamClient.getBearerToken(request)
}

func (iamClient *IamClient) getBearerToken(requestBody map[string]string) (*model.ClientTokenResponse, error) {
	headers := map[string]string{
		"Accept":       "*/*",
		"Content-Type": "x-www-form-urlencoded",
	}

	httpRequest, err := iamClient.client.CreateWithParams(POST, iamClient.iamHost, iamClient.tokenApi, headers, nil, requestBody)
	if err != nil {
		return nil, err
	}

	httpResponse, err := iamClient.client.Send(httpRequest)
	if err != nil {
		return nil, err
	}

	parsedResponse, err := httpclient.Parse(httpResponse, &model.ClientTokenResponse{})
	if err != nil {
		return nil, err
	}
	return parsedResponse, nil
}

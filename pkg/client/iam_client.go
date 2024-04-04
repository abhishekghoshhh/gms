package client

import (
	"errors"
	"github.com/abhishekghoshhh/gms/pkg/httpclient"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"net/http"
	"strconv"
)

const (
	MediaTypeAll          = "*/*"
	ApplicationUrlEncoded = "x-www-form-urlencoded"
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
		"Accept":        MediaTypeAll,
	}

	response, err := iamClient.client.MakeRequest(http.MethodGet, iamClient.iamHost, iamClient.scimProfileApi, headers, nil, nil)

	if err != nil {
		return nil, err
	}
	return httpclient.Parse(response, &model.IamProfileResponse{})
}

func (iamClient *IamClient) FetchUserCount(token string) (*model.IamProfileListResponse, error) {
	headers := map[string]string{
		"Authorization": token,
		"Accept":        MediaTypeAll,
	}
	response, err := iamClient.client.MakeRequest(http.MethodGet, iamClient.iamHost, iamClient.getUserCountApi, headers, nil, nil)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return httpclient.Parse(response, &model.IamProfileListResponse{})
}

func (iamClient *IamClient) FetchUsersInBatch(token string, startingIndex, count int) (*model.IamProfileListResponse, error) {
	headers := map[string]string{
		"Authorization": token,
		"Accept":        MediaTypeAll,
	}

	queryParams := map[string]string{
		"startIndex": strconv.Itoa(startingIndex),
		"count":      strconv.Itoa(count),
	}
	response, err := iamClient.client.MakeRequest(http.MethodGet, iamClient.iamHost, iamClient.fetchUsersInBatchApi, headers, queryParams, nil)

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
		"Accept":       MediaTypeAll,
		"Content-Type": ApplicationUrlEncoded,
	}

	httpResponse, err := iamClient.client.MakeRequest(http.MethodPost, iamClient.iamHost, iamClient.tokenApi, headers, nil, requestBody)
	if err != nil {
		return nil, err
	}

	parsedResponse, err := httpclient.Parse(httpResponse, &model.ClientTokenResponse{})
	if err != nil {
		return nil, err
	}
	return parsedResponse, nil
}

func (iamClient *IamClient) FetchUserByCertSubject(token, subject string) (*model.IamProfileListResponse, error) {
	queryParams := map[string]string{
		"certificateSubject": subject,
	}

	headers := map[string]string{
		"Accept":        MediaTypeAll,
		"Authorization": token,
	}

	response, err := iamClient.client.MakeRequest(http.MethodGet, iamClient.iamHost, iamClient.findAccountByCertSubjectApi, headers, queryParams, nil)
	if err != nil {
		return nil, err
	}

	parsedResponse, err := httpclient.Parse(response, &model.IamProfileListResponse{})
	if err != nil {
		return nil, err
	}
	return parsedResponse, nil
}

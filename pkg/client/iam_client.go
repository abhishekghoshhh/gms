package client

import (
	"errors"
	"net/http"
	"strconv"

	httpclient "github.com/abhishekghoshhh/gms/pkg/http"
	"github.com/abhishekghoshhh/gms/pkg/model"
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
	fetchUserByIdApi            string
	client                      *httpclient.CustomClient
}

func New(client *httpclient.CustomClient,
	iamHost,
	scimProfileApi, tokenApi, findAccountByCertSubjectApi, getUserCountApi, fetchUsersInBatchApi, fetchUserByIdApi string) (*IamClient, error) {

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
		fetchUserByIdApi,
		client,
	}, nil
}

func (iamClient *IamClient) FetchUser(token string) (*model.IamProfileResponse, error) {
	headers := map[string]string{
		"Authorization": token,
		"Accept":        MediaTypeAll,
	}
	req := httpclient.
		Request(
			iamClient.iamHost,
			iamClient.scimProfileApi,
			http.MethodGet,
		).
		Headers(headers).
		Timeout(2) //timeout needs to be changed

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.IamProfileResponse{})
	}
}

func (iamClient *IamClient) FetchUserCount(token string) (*model.IamProfileListResponse, error) {
	headers := map[string]string{
		"Authorization": token,
		"Accept":        MediaTypeAll,
	}

	req := httpclient.
		Request(
			iamClient.iamHost,
			iamClient.getUserCountApi,
			http.MethodGet,
		).
		Headers(headers).
		Timeout(2) //timeout needs to be changed

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.IamProfileListResponse{})
	}
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

	req := httpclient.
		Request(
			iamClient.iamHost,
			iamClient.fetchUsersInBatchApi,
			http.MethodGet,
		).
		Headers(headers).
		QueryParams(queryParams).
		Timeout(2) //timeout needs to be changed

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.IamProfileListResponse{})
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

func (iamClient *IamClient) FetchUserByCertSubject(token, subject string) (*model.IamProfileListResponse, error) {
	queryParams := map[string]string{
		"certificateSubject": subject,
	}

	headers := map[string]string{
		"Accept":        MediaTypeAll,
		"Authorization": token,
	}

	req := httpclient.
		Request(
			iamClient.iamHost,
			iamClient.findAccountByCertSubjectApi,
			http.MethodGet,
		).
		Headers(headers).
		QueryParams(queryParams).
		Timeout(2) //timeout needs to be changed

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.IamProfileListResponse{})
	}
}

func (iamClient *IamClient) FetchUserById(token, userId string) (*model.IamProfileListResponse, error) {
	queryParams := map[string]string{
		"userId": userId,
	}

	headers := map[string]string{
		"Accept":        MediaTypeAll,
		"Authorization": token,
	}

	req := httpclient.
		Request(
			iamClient.iamHost,
			iamClient.findAccountByCertSubjectApi,
			http.MethodGet,
		).
		Headers(headers).
		QueryParams(queryParams).
		Timeout(2) //timeout needs to be changed

	if resp, err := iamClient.client.Send(req); err != nil {
		return nil, err
	} else {
		return httpclient.Parse(resp, &model.IamProfileListResponse{})
	}
}

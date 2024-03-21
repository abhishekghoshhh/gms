package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/abhishekghoshhh/gms/pkg/logger"
	"github.com/abhishekghoshhh/gms/pkg/model"
	"go.uber.org/zap"
)

type IamClient struct {
	iamHost        string
	scimProfileApi string
}

func New(iamHost, scimProfileApi string) *IamClient {
	return &IamClient{
		iamHost,
		scimProfileApi,
	}
}

func (iamClient *IamClient) FetchUser(token string) (*model.IamProfileResponse, error) {
	iamUrl, err := url.Parse(iamClient.iamHost)
	if err != nil {
		log.Fatal("Error creating request:", err)
		return nil, err
	}
	iamUrl.Path = iamClient.scimProfileApi
	req, err := http.NewRequest("GET", iamUrl.String(), bytes.NewBuffer(nil))
	if err != nil {
		log.Fatal("Error creating request:", err)
		return nil, err
	}
	logger.Info("url is " + (iamClient.iamHost + iamClient.scimProfileApi))
	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on response:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
		return nil, err
	}
	var iamProfileResponse model.IamProfileResponse
	if err := json.Unmarshal(body, &iamProfileResponse); err != nil {
		logger.Error("error is " + err.Error())
		return nil, errors.New("invalid profileResponse")
	}
	logger.Info("profile response is ", zap.Any("resp", iamProfileResponse))
	return &iamProfileResponse, nil
}

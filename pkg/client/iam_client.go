package client

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/abhishekghoshhh/gms/pkg/logger"
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

func (iamClient *IamClient) FetchUser(token string) (string, error) {
	req, err := http.NewRequest("GET", iamClient.iamHost+iamClient.scimProfileApi, bytes.NewBuffer(nil))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	logger.Info("url is " + (iamClient.iamHost + iamClient.scimProfileApi))
	logger.Info("token is " + token)
	req.Header.Set("Authorization", token)
	req.Header.Add("Accept", "*/*")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on response:", err)
	}
	defer resp.Body.Close()
	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	logger.Info("response is " + string(body))
	return string(body), nil
}

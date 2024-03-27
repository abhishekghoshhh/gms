package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/abhishekghoshhh/gms/pkg/logger"
	"go.uber.org/zap"
)

type Client struct {
	client *http.Client
}

func NewClient(client *http.Client) *Client {
	return &Client{
		client,
	}
}

func (c *Client) Send(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		logger.Error("Error on response" + err.Error())
		return nil, err
	}
	readingErr := resp.Body.Close()
	if readingErr != nil {
		return nil, readingErr
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logger.Error("Error reading response body:" + err.Error())
		return nil, err
	}
	return body, nil
}

func (*Client) Create(method, host, path string, headers map[string]string) (*http.Request, error) {
	parsedUrl, err := url.Parse(host)
	if err != nil {
		logger.Error("Error constructing the url" + err.Error())
		return nil, err
	}
	parsedUrl.Path = path

	req, err := http.NewRequest(method, parsedUrl.String(), bytes.NewBuffer(nil))

	if err != nil {
		logger.Error("Error creating new request" + err.Error())
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

func (*Client) CreateWithBody(method, host, path string, headers map[string]string, body ...any) (*http.Request, error) {
	var err error
	parsedUrl, err := url.Parse(host)
	if err != nil {
		logger.Error("Error constructing the url" + err.Error())
		return nil, err
	}
	parsedUrl.Path = path

	var req *http.Request
	var reqBody []byte = nil

	if len(body) != 0 {
		reqBody, err = json.Marshal(body[0])
	}
	if err != nil {
		logger.Error("Error marshaling request body" + err.Error())
		return nil, err
	}
	req, err = http.NewRequest(method, parsedUrl.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		logger.Error("Error creating new request" + err.Error())
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

func Parse[T any](data []byte, dataObject *T) (*T, error) {
	if err := json.Unmarshal(data, dataObject); err != nil {
		logger.Error("error is " + err.Error())
		return nil, errors.New("invalid response")
	}
	logger.Debug("response is", zap.Any("resp", dataObject))
	return dataObject, nil
}

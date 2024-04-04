package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/abhishekghoshhh/gms/pkg/logger"
	"go.uber.org/zap"
)

type Client struct {
	client  *http.Client
	timeout time.Duration
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		client:  http.DefaultClient,
		timeout: timeout,
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

func (*Client) createUrl(host, path string, queryParams map[string]string) (*url.URL, error) {
	newUrl, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	newUrl.Path = path
	if queryParams != nil {
		queries := newUrl.Query()
		for key, val := range queryParams {
			queries.Add(key, val)
		}
		newUrl.RawQuery = queries.Encode()
	}
	return newUrl, nil
}

func (c *Client) Create(method, host, path string, headers map[string]string) (*http.Request, error) {
	parsedUrl, err := c.createUrl(host, path, nil)
	if err != nil {
		logger.Error("Error constructing the url" + err.Error())
		return nil, err
	}

	req, err := c.createRequest(method, parsedUrl.String(), bytes.NewBuffer(nil))

	if err != nil {
		logger.Error("Error creating new request" + err.Error())
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

func (c *Client) createRequest(method, url string, body *bytes.Buffer) (*http.Request, error) {
	deadline := time.Now().Add(c.timeout * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	return req, err
}

func (c *Client) CreateWithParams(method, host, path string, headers map[string]string, queryParams map[string]string, body any) (*http.Request, error) {
	var err error

	parsedUrl, err := c.createUrl(host, path, queryParams)
	if err != nil {
		logger.Error("Error constructing the url" + err.Error())
		return nil, err
	}

	var req *http.Request
	var reqBody []byte = nil

	if body != nil {
		reqBody, err = json.Marshal(body)
	}
	if err != nil {
		logger.Error("Error marshaling request body" + err.Error())
		return nil, err
	}
	req, err = c.createRequest(method, parsedUrl.String(), bytes.NewBuffer(reqBody))
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

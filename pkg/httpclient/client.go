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

func (c *Client) send(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		if errors.Is(req.Context().Err(), context.DeadlineExceeded) {
			logger.Error("Deadline exceeded, request failed")
			return nil, req.Context().Err()
		}
		logger.Error("Error on response " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		logger.Error("Error reading response body:" + err.Error())
		return nil, err
	}
	return body, nil
}

func (c *Client) createRequest(method, url string, body *bytes.Buffer) (*http.Request, error) {
	deadline := time.Now().Add(c.timeout)

	_, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	req, err := http.NewRequest(method, url, body)
	return req, err
}

func (c *Client) MakeRequest(method, host, path string, headers map[string]string, queryParams map[string]string, body any) ([]byte, error) {
	parsedUrl, err := url.Parse(host)
	if err != nil {
		logger.Error("Error constructing the url" + err.Error())
		return nil, err
	}

	parsedUrl.Path = path
	if queryParams != nil {
		queries := parsedUrl.Query()
		for key, val := range queryParams {
			queries.Add(key, val)
		}
		parsedUrl.RawQuery = queries.Encode()
	}

	var reqBody []byte = nil

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			logger.Error("Error marshaling request body" + err.Error())
			return nil, err
		}
	}

	logger.Debug("calling " + parsedUrl.String())

	req, err := c.createRequest(method, parsedUrl.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		logger.Error("Error creating new request" + err.Error())
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return c.send(req)
}

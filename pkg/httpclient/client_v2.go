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

var (
	client = http.DefaultClient
)

const (
	MaxTimeout = 300 * time.Second
)

func Send[T any](config *RequestConfiguration, responseType *T) (*T, error) {
	parsedUrl, err := config.prepareUrl()
	if err != nil {
		return nil, err
	}
	logger.Debug("calling " + parsedUrl.String())

	reqBody, err := config.prepareBody()
	if err != nil {
		return nil, err
	}

	var req *http.Request

	ctx, cancel := prepareContext(config.timeout)
	defer cancel()
	req, err = http.NewRequestWithContext(ctx, config.method, parsedUrl.String(), bytes.NewBuffer(reqBody))

	if err != nil {
		logger.Error("Error creating new request" + err.Error())
		return nil, err
	}

	for key, value := range config.headers {
		req.Header.Set(key, value)
	}

	resp, err := send(client, req)
	if err != nil {
		return nil, err
	}
	return Parse(resp, responseType)
}

func send(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
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
func Parse[T any](data []byte, dataObject *T) (*T, error) {
	if err := json.Unmarshal(data, dataObject); err != nil {
		logger.Error("error is " + err.Error())
		return nil, errors.New("invalid response")
	}
	logger.Debug("response is", zap.Any("resp", dataObject))
	return dataObject, nil
}

func prepareContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	if timeout == 0 {
		timeout = MaxTimeout
	}
	deadline := time.Now().Add(timeout)
	return context.WithDeadline(context.Background(), deadline)
}

type RequestConfiguration struct {
	host        string
	path        string
	method      string
	headers     map[string]string
	queryParams map[string]string
	body        any
	timeout     time.Duration
}

func Request(host, path, method string) *RequestConfiguration {
	return &RequestConfiguration{
		host:   host,
		path:   path,
		method: method,
	}
}
func (config *RequestConfiguration) Headers(headers map[string]string) *RequestConfiguration {
	config.headers = headers
	return config
}
func (config *RequestConfiguration) QueryParams(queryParams map[string]string) *RequestConfiguration {
	config.queryParams = queryParams
	return config
}
func (config *RequestConfiguration) Body(body any) *RequestConfiguration {
	config.body = body
	return config
}
func (config *RequestConfiguration) Timeout(timeout int) *RequestConfiguration {
	config.timeout = time.Duration(timeout) * time.Second
	return config
}

func (config *RequestConfiguration) prepareUrl() (*url.URL, error) {
	newUrl, err := url.Parse(config.host)
	if err != nil {
		logger.Error("Error constructing the newUrl " + err.Error())
		return nil, err
	}
	newUrl.Path = config.path
	if config.queryParams != nil {
		queries := newUrl.Query()
		for key, val := range config.queryParams {
			queries.Add(key, val)
		}
		newUrl.RawQuery = queries.Encode()
	}
	return newUrl, nil
}

func (config *RequestConfiguration) prepareBody() ([]byte, error) {
	if config.body == nil {
		return nil, nil
	}
	reqBody, err := json.Marshal(config.body)
	if err != nil {
		logger.Error("Error marshaling request body " + err.Error())
		return nil, err
	}
	return reqBody, nil
}

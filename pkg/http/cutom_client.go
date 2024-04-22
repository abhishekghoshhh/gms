	package http

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
	MAX_TIMEOUT = 300 * time.Second
)

type HttpClient interface {
	Send(config *RequestConf) ([]byte, error)
}

type CustomClient struct {
}

func NewClient() *CustomClient {
	return &CustomClient{}
}

func (CustomClient) Send(conf *RequestConf) ([]byte, error) {
	url, err := conf.prepareUrl()
	if err != nil {
		return nil, err
	}
	logger.Debug("calling " + url.String())

	reqBody, err := conf.prerpareBody()
	if err != nil {
		return nil, err
	}

	var req *http.Request

	ctx, cancel := prepareContext(conf.timeout)
	defer cancel()
	req, err = http.NewRequestWithContext(ctx, conf.method, url.String(), bytes.NewBuffer(reqBody))

	if err != nil {
		logger.Error("Error creating new request" + err.Error())
		return nil, err
	}

	for key, value := range conf.headers {
		req.Header.Set(key, value)
	}

	return send(client, req)
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
		timeout = MAX_TIMEOUT
	}
	deadline := time.Now().Add(timeout)
	return context.WithDeadline(context.Background(), deadline)
}

type RequestConf struct {
	host        string
	path        string
	method      string
	headers     map[string]string
	queryParams map[string]string
	body        any
	timeout     time.Duration
}

func Request(host, path, method string) *RequestConf {
	return &RequestConf{
		host:   host,
		path:   path,
		method: method,
	}
}
func (config *RequestConf) Headers(headers map[string]string) *RequestConf {
	config.headers = headers
	return config
}
func (config *RequestConf) QueryParams(queryParams map[string]string) *RequestConf {
	config.queryParams = queryParams
	return config
}
func (config *RequestConf) Body(body any) *RequestConf {
	config.body = body
	return config
}
func (config *RequestConf) Timeout(timeout int) *RequestConf {
	config.timeout = time.Duration(timeout) * time.Second
	return config
}

func (config *RequestConf) prepareUrl() (*url.URL, error) {
	url, err := url.Parse(config.host)
	if err != nil {
		logger.Error("Error constructing the url " + err.Error())
		return nil, err
	}
	url.Path = config.path
	if config.queryParams != nil {
		queries := url.Query()
		for key, val := range config.queryParams {
			queries.Add(key, val)
		}
		url.RawQuery = queries.Encode()
	}
	return url, nil
}

func (config *RequestConf) prerpareBody() ([]byte, error) {
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

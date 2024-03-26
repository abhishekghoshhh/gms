package httpclient

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	client         *http.Client
	requestCreator *http.Request
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c *Client) Send(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal("Error on response", err)
		return nil, err
	}
	readingErr := resp.Body.Close()
	if readingErr != nil {
		return nil, readingErr
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading response body:", err)
		return nil, err
	}
	return body, nil
}

func (*Client) Create(method, host, path string, headers map[string]string) (*http.Request, error) {
	parsedUrl, err := url.Parse(host)
	parsedUrl.Path = path

	req, err := http.NewRequest(method, parsedUrl.String(), bytes.NewBuffer(nil))

	if err != nil {
		log.Fatal("Error creating new request", err)
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

// Package httpclient is the http client wrapper to make http requests to server
// Copyright 2024 The Oleg Nazarov. All rights reserved.
//
// This package http client for agent.
package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/ole-larsen/plutonium/internal/compressor"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/ole-larsen/plutonium/models"
)

var logger = log.NewLogger("info", log.DefaultBuildLogger)

// HTTPClient - custom agent wrapper for http.Client.
type HTTPClient struct {
	Client   *http.Client
	settings *settings.Settings
}

var (
	singleton *HTTPClient
	once      sync.Once
)

func NewHTTPClient() *HTTPClient {
	once.Do(func() {
		singleton = &HTTPClient{
			Client: &http.Client{
				Timeout: time.Minute,
			},
		}
	})

	return singleton
}

func SetDefaultTransport() *http.Transport {
	if transport, ok := http.DefaultTransport.(*http.Transport); ok {
		return transport.Clone()
	}

	return nil
}

func (c *HTTPClient) SetTransport(t http.RoundTripper) *HTTPClient {
	c.Client.Transport = t
	return c
}

func (c *HTTPClient) SetTimeout(t time.Duration) *HTTPClient {
	c.Client.Timeout = t
	return c
}

func (c *HTTPClient) SetSettings(cfg *settings.Settings) *HTTPClient {
	c.settings = cfg
	return c
}

func (c *HTTPClient) GetTransport() http.RoundTripper {
	return c.Client.Transport
}

func (c *HTTPClient) GetSettings() *settings.Settings {
	return c.settings
}

func (c *HTTPClient) Get(url string) (resp *http.Response, err error) {
	return c.Client.Get(url)
}

func (c *HTTPClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return c.Client.Post(url, contentType, body)
}

func (c *HTTPClient) SetRequest(method, url string, body []byte) (*http.Request, error) {
	var (
		request *http.Request
		err     error
	)

	if len(body) > 0 {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	}

	if len(body) == 0 {
		request, err = http.NewRequest(method, url, http.NoBody)
	}

	return request, err
}

func (c *HTTPClient) SetRequestWithContext(ctx context.Context, method, url string, body []byte) (*http.Request, error) {
	var (
		request *http.Request
		err     error
	)

	if len(body) > 0 {
		request, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	}

	if len(body) == 0 {
		request, err = http.NewRequestWithContext(ctx, method, url, http.NoBody)
	}

	return request, err
}

func (c *HTTPClient) SetHeaders(request *http.Request, headers map[string]string) *http.Request {
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	return request
}

func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

func (c *HTTPClient) MakeRequest(
	ctx context.Context,
	method string,
	url string,
	body []byte,
) ([]byte, error) {
	var err error

	request, err := c.SetRequestWithContext(ctx, method, url, body)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}

	resp, err := c.Do(request)

	if err != nil {
		return nil, err
	}

	var data []byte

	if resp.Header.Get("Acccept-Encoding") == "gzip" {
		var cr *compressor.CompressReader

		cr, err = compressor.NewCompressReader(resp.Body)
		if err != nil {
			logger.Errorln(err)
			return nil, NewError(err)
		}

		resp.Body = cr

		defer func() {
			e := cr.Close()
			if e != nil {
				logger.Errorln(e)
			}
		}()
		defer func() {
			e := resp.Body.Close()
			if e != nil {
				logger.Errorln(e)
			}
		}()
	} else {
		data, err = io.ReadAll(resp.Body)

		if err != nil {
			logger.Errorln(err)
			return nil, NewError(fmt.Errorf("error reading response body: %w", err))
		}
	}

	if resp.StatusCode != http.StatusOK {
		logger.Infoln(url, resp.StatusCode)
	}

	switch status := resp.StatusCode; {
	case status >= 200 && status < 300: // 2xx
		return data, nil
	case status >= 400 && status < 500: // 4xx
		return nil, NewError(fmt.Errorf("error response: %s", string(data)))
	default: // 5xx
		return nil, NewError(fmt.Errorf("error response: %s", string(data)))
	}
}

func (c *HTTPClient) GetCredentials(ctx context.Context, clientID string) (*models.Credentials, error) {
	url := c.settings.OAUTH2.Provider + "/api/v1/credentials?client_id=" + clientID

	body, err := c.MakeRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.Credentials
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *HTTPClient) Authorize(ctx context.Context, authorizeURL string) (*models.Callback, error) {
	body, err := c.MakeRequest(ctx, "GET", authorizeURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.Callback
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

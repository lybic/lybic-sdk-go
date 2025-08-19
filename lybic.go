package lybic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	envOrgId    = "LYBIC_ORG_ID"
	envApiKey   = "LYBIC_API_KEY"
	envEndpoint = "LYBIC_API_ENDPOINT"

	defaultEndpoint = "https://api.lybic.cn"
	defaultTimeout  = 10 // seconds
)

var (
	ErrNeedOrgId    = errors.New("please specify an organization")
	ErrNeedEndpoint = errors.New("please specify an API endpoint")
)

type Config struct {
	OrgId    string // Organization ID, required for API access
	ApiKey   string // API Key for authentication (Not necessary)
	Endpoint string // API endpoint, default is "https://api.lybic.cn"
	Timeout  uint8  // Timeout in seconds for HTTP requests, default is 10 seconds

	ExtraHeaders map[string]string // Additional headers to be sent with each request

	Logger Logger // Logger interface for logging, can be nil for no logging
}

func NewConfig() *Config {
	return &Config{
		OrgId:    getEnv(envOrgId, ""),
		ApiKey:   getEnv(envApiKey, ""),
		Endpoint: getEnv(envEndpoint, defaultEndpoint),

		Timeout: defaultTimeout,
	}
}

type Client struct {
	client *http.Client

	config *Config
}

func (c *Client) GetConfig() *Config {
	return c.config
}

func NewClient(config *Config) (*Client, error) {
	if config == nil {
		config = NewConfig()
	}

	if config.Logger == nil {
		config.Logger = &emptyLogger{}
	}

	if config.Endpoint == "" {
		config.Logger.Errorf("API endpoint is not set, please specify it in config or set the %s environment variable", envEndpoint)
		return nil, ErrNeedEndpoint
	}
	if config.OrgId == "" {
		config.Logger.Errorf("organization id is not set, please specify it in config or set the %s environment variable", envOrgId)
		return nil, ErrNeedOrgId
	}

	// Remove trailing slash from endpoint
	config.Endpoint = strings.TrimSuffix(config.Endpoint, "/")

	return &Client{
		client: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
		config: config,
	}, nil
}

func (c *Client) request(ctx context.Context, method, url string, params map[string]string, bodyDto any) (*http.Response, error) {
	var body io.Reader
	if bodyDto != nil {
		data, err := json.Marshal(bodyDto)
		if err != nil {
			c.config.Logger.Errorf("failed to marshal request body: %v", err)
			return nil, err
		}

		c.config.Logger.Debugf("request body: %s", string(data))
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.config.Endpoint+url, body)
	if err != nil {
		c.config.Logger.Errorf("failed to create request: %v", err)
		return nil, err
	}

	// Set common headers
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.config.ApiKey != "" {
		req.Header.Set("x-api-key", c.config.ApiKey)
	}
	// Add extra headers if any
	for k, v := range c.config.ExtraHeaders {
		c.config.Logger.Debugf("Setting header `%s: %s`", k, v)
		req.Header.Set(k, v)
	}

	// Add query parameters
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return c.client.Do(req)
}

func getEnv(key string, defaultVal string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultVal
}

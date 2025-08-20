// Copyright (c) 2019-2025   Beijing Tingyu Technology Co., Ltd.
// Copyright (c) 2025        Lybic Development Team <team@lybic.ai, lybic@tingyutech.com>
// Copyright (c) 2025        Lu Yicheng <luyicheng@tingyutech.com>
//
// These Terms of Service ("Terms") set forth the rules governing your access to and use of the website lybic.ai
// ("Website"), our web applications, and other services (collectively, the "Services") provided by Beijing Tingyu
// Technology Co., Ltd. ("Company," "we," "us," or "our"), a company registered in Haidian District, Beijing. Any
// breach of these Terms may result in the suspension or termination of your access to the Services.
// By accessing and using the Services and/or the Website, you represent that you are at least 18 years old,
// acknowledge that you have read and understood these Terms, and agree to be bound by them. By using or accessing
// the Services and/or the Website, you further represent and warrant that you have the legal capacity and authority
// to agree to these Terms, whether as an individual or on behalf of a company. If you do not agree to all of these
// Terms, do not access or use the Website or Services.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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

type client struct {
	client *http.Client
	config *Config
}

func (c *client) GetConfig() *Config {
	return c.config
}

// headerTransport is custom transport to add headers to all requests.
type headerTransport struct {
	base    http.RoundTripper
	headers map[string]string
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := *req
	newReq.Header = req.Header.Clone()
	for key, value := range t.headers {
		newReq.Header.Set(key, value)
	}
	return t.base.RoundTrip(&newReq)
}

func newClient(config *Config) (*client, error) {
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

	// Prepare headers for the custom transport
	headers := make(map[string]string)
	if config.ApiKey != "" {
		headers["x-api-key"] = config.ApiKey
	}
	for k, v := range config.ExtraHeaders {
		config.Logger.Debugf("Setting persistent header `%s: %s`", k, v)
		headers[k] = v
	}

	var baseTransport http.RoundTripper
	if config.HttpTransport != nil {
		baseTransport = config.HttpTransport
	} else {
		baseTransport = http.DefaultTransport
	}

	var transport http.RoundTripper
	if len(headers) > 0 {
		transport = &headerTransport{
			base:    baseTransport,
			headers: headers,
		}
	}

	return &client{
		client: &http.Client{
			Timeout:   time.Duration(config.Timeout) * time.Second,
			Transport: transport,
		},
		config: config,
	}, nil
}

func (c *client) request(ctx context.Context, method, url string, params map[string]string, bodyDto any) (*http.Response, error) {
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
	// ApiKey and ExtraHeaders are now handled by the custom transport.

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

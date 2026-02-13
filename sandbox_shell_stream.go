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
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lybic/lybic-sdk-go/pkg/json"
)

// CreateSandboxShellCommandStream creates a shell session with streaming output (SSE).
// It returns a channel that emits shell stream events and an error.
// The channel will be closed when the stream ends or an error occurs.
func (c *client) CreateSandboxShellCommandStream(ctx context.Context, sandboxId string,
	dto SandboxShellCommandStreamCreateRequestDto) (<-chan SandboxShellStreamEvent, error) {
	c.config.Logger.Info("Creating sandbox shell command stream", "sandboxId:", sandboxId)

	url := fmt.Sprintf("%s/api/orgs/%s/sandboxes/%s/shell/stream", c.config.Endpoint, c.config.OrgId, sandboxId)

	var body io.Reader
	if dto.Command != "" {
		data, err := json.Marshal(dto)
		if err != nil {
			c.config.Logger.Errorf("failed to marshal request body: %v", err)
			return nil, err
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		c.config.Logger.Errorf("failed to create request: %v", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := c.client.Do(req)
	if err != nil {
		c.config.Logger.Errorf("failed to execute request: %v", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	eventChan := make(chan SandboxShellStreamEvent, 10)

	go func() {
		defer close(eventChan)
		defer resp.Body.Close()

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					c.config.Logger.Errorf("error reading stream: %v", err)
				}
				return
			}

			line = bytes.TrimSpace(line)
			if len(line) == 0 {
				continue
			}

			if !bytes.HasPrefix(line, []byte("data: ")) {
				continue
			}

			data := bytes.TrimPrefix(line, []byte("data: "))
			event, err := parseSSEEvent(data)
			if err != nil {
				c.config.Logger.Errorf("error parsing SSE event: %v", err)
				continue
			}

			select {
			case eventChan <- event:
				if event.Type == "end" || event.Type == "waiting" {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventChan, nil
}

func parseSSEEvent(data []byte) (SandboxShellStreamEvent, error) {
	var eventData map[string]string
	if err := json.Unmarshal(data, &eventData); err != nil {
		return SandboxShellStreamEvent{}, err
	}

	for eventType, encodedData := range eventData {
		decodedData := ""
		if encodedData != "" {
			decoded, err := base64.StdEncoding.DecodeString(encodedData)
			if err != nil {
				return SandboxShellStreamEvent{}, fmt.Errorf("failed to decode base64 data: %w", err)
			}
			decodedData = string(decoded)
		}

		return SandboxShellStreamEvent{
			Type: strings.TrimSpace(eventType),
			Data: decodedData,
		}, nil
	}

	return SandboxShellStreamEvent{}, fmt.Errorf("no event data found")
}

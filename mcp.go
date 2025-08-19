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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ListMcpServers returns a list of MCP servers for the organization.
func (c *client) ListMcpServers(ctx context.Context) ([]McpServerResponseDto, error) {
	c.config.Logger.Info("Listing mcp servers", "servers:", c.config.OrgId)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list mcp servers: %s", resp.Status)
	}

	var mcpServers []McpServerResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&mcpServers); err != nil {
		return nil, err
	}

	return mcpServers, nil
}

// CreateMcpServer creates a new MCP server.
func (c *client) CreateMcpServer(ctx context.Context, dto CreateMcpServerDto) (*McpServerResponseDto, error) {
	c.config.Logger.Info("Creating mcp server", "dto:", dto)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create mcp server: %s", resp.Status)
	}

	var mcpServer McpServerResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&mcpServer); err != nil {
		return nil, err
	}

	return &mcpServer, nil
}

// GetDefaultMcpServer returns the default MCP server for the organization.
func (c *client) GetDefaultMcpServer(ctx context.Context) (*McpServerResponseDto, error) {
	c.config.Logger.Info("Getting default mcp server", "servers:", c.config.OrgId)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers/default", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get default mcp server: %s", resp.Status)
	}

	var mcpServer McpServerResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&mcpServer); err != nil {
		return nil, err
	}

	return &mcpServer, nil
}

// DeleteMcpServer deletes an MCP server by its ID.
func (c *client) DeleteMcpServer(ctx context.Context, mcpServerId string) error {
	c.config.Logger.Info("Deleting mcp server", "server:", mcpServerId)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers/%s", c.config.OrgId, mcpServerId)
	resp, err := c.request(ctx, http.MethodDelete, url, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete mcp server: %s", resp.Status)
	}

	return nil
}

// SetMcpServerToSandbox sets the specified MCP server to the given Sandbox.
func (c *client) SetMcpServerToSandbox(ctx context.Context, mcpServerId string, dto SetMcpServerToSandboxResponseDto) error {
	c.config.Logger.Info("Setting mcp server to sandbox", "server:", mcpServerId)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers/%s/sandbox", c.config.OrgId, mcpServerId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to set mcp server to sandbox: %s", resp.Status)
	}

	return nil
}

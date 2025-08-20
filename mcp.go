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

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ListMcpServers returns a list of MCP servers for the organization.
func (m *mcpClient) ListMcpServers(ctx context.Context) ([]McpServerResponseDto, error) {
	m.client.config.Logger.Info("Listing mcp servers", "servers:", m.client.config.OrgId)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers", m.client.config.OrgId)
	resp, err := m.client.request(ctx, http.MethodGet, url, nil, nil)
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
func (m *mcpClient) CreateMcpServer(ctx context.Context, dto CreateMcpServerDto) (*McpServerResponseDto, error) {
	m.client.config.Logger.Info("Creating mcp server", "dto:", dto)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers", m.client.config.OrgId)
	resp, err := m.client.request(ctx, http.MethodPost, url, nil, dto)
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
func (m *mcpClient) GetDefaultMcpServer(ctx context.Context) (*McpServerResponseDto, error) {
	m.client.config.Logger.Info("Getting default mcp server", "servers:", m.client.config.OrgId)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers/default", m.client.config.OrgId)
	resp, err := m.client.request(ctx, http.MethodGet, url, nil, nil)
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
func (m *mcpClient) DeleteMcpServer(ctx context.Context, mcpServerId string) error {
	m.client.config.Logger.Info("Deleting mcp server", "server:", mcpServerId)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers/%s", m.client.config.OrgId, mcpServerId)
	resp, err := m.client.request(ctx, http.MethodDelete, url, nil, nil)
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
func (m *mcpClient) SetMcpServerToSandbox(ctx context.Context, mcpServerId string, dto SetMcpServerToSandboxResponseDto) error {
	m.client.config.Logger.Info("Setting mcp server to sandbox", "server:", mcpServerId)
	url := fmt.Sprintf("/api/orgs/%s/mcp-servers/%s/sandbox", m.client.config.OrgId, mcpServerId)
	resp, err := m.client.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to set mcp server to sandbox: %s", resp.Status)
	}

	return nil
}

type mcpClient struct {
	session *mcp.ClientSession
	client  *client
}

func newMcpClient(ctx context.Context, client *client, address *string) (*mcpClient, error) {
	m := &mcpClient{client: client}

	var serverAddress string
	if address == nil {
		// use default server
		mcpServer, err := m.GetDefaultMcpServer(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get default MCP server: %w", err)
		}
		serverAddress = fmt.Sprintf("%s/api/mcp/%s", m.client.config.Endpoint, mcpServer.Id)
	} else {
		serverAddress = fmt.Sprintf("%s/api/mcp/%s", m.client.config.Endpoint, *address)
		m.client.config.Logger.Infof("Using specific MCP server address: %s", serverAddress)
	}

	cli := mcp.NewClient(&mcp.Implementation{Name: "mcp-client/lybic-sdk-go", Version: "v0.0.2"}, nil)
	transport := mcp.NewStreamableClientTransport(serverAddress, &mcp.StreamableClientTransportOptions{HTTPClient: client.client})

	session, err := cli.Connect(ctx, transport)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MCP server: %w", err)
	}

	m.session = session
	return m, nil
}

func (m *mcpClient) Close() error {
	return m.session.Close()
}

func (m *mcpClient) CallTools(ctx context.Context, args map[string]any, service *string) (*mcp.CallToolResult, error) {
	if service == nil {
		service = new(string)
		*service = "computer-use"
	}

	response, err := m.session.CallTool(ctx, &mcp.CallToolParams{
		Name:      *service,
		Arguments: args,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call tools: %w", err)
	}
	return response, nil
}

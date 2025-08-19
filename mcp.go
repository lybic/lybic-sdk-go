package lybic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ListMcpServers returns a list of MCP servers for the organization.
func (c *Client) ListMcpServers(ctx context.Context) ([]McpServerResponseDto, error) {
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
func (c *Client) CreateMcpServer(ctx context.Context, dto CreateMcpServerDto) (*McpServerResponseDto, error) {
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
func (c *Client) GetDefaultMcpServer(ctx context.Context) (*McpServerResponseDto, error) {
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
func (c *Client) DeleteMcpServer(ctx context.Context, mcpServerId string) error {
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
func (c *Client) SetMcpServerToSandbox(ctx context.Context, mcpServerId string, dto SetMcpServerToSandboxResponseDto) error {
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

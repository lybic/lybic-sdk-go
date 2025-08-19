package lybic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ListSandboxes returns a list of sandboxes for the organization.
func (c *Client) ListSandboxes(ctx context.Context) ([]GetSandboxResponseDtoSandbox, error) {
	c.config.Logger.Info("Listing sandboxes")

	url := fmt.Sprintf("/api/orgs/%s/sandboxes", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list sandboxes: %s", resp.Status)
	}

	var sandboxes []GetSandboxResponseDtoSandbox
	if err := json.NewDecoder(resp.Body).Decode(&sandboxes); err != nil {
		return nil, err
	}

	return sandboxes, nil
}

// CreateSandbox creates a new sandbox.
func (c *Client) CreateSandbox(ctx context.Context, dto CreateSandboxDto) (*GetSandboxResponseDto, error) {
	c.config.Logger.Info("Creating sandbox", "dto:", dto)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create sandbox: %s", resp.Status)
	}

	var sandbox GetSandboxResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&sandbox); err != nil {
		return nil, err
	}

	return &sandbox, nil
}

// GetSandbox retrieves the details of a sandbox by its ID.
func (c *Client) GetSandbox(ctx context.Context, sandboxId string) (*GetSandboxResponseDto, error) {
	c.config.Logger.Info("Getting sandbox info", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get sandbox: %s", resp.Status)
	}

	var sandbox GetSandboxResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&sandbox); err != nil {
		return nil, err
	}

	return &sandbox, nil
}

// DeleteSandbox deletes a sandbox by its ID.
func (c *Client) DeleteSandbox(ctx context.Context, sandboxId string) error {
	c.config.Logger.Info("Deleting sandbox", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodDelete, url, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete sandbox: %s", resp.Status)
	}

	return nil
}

// ExtendSandbox extends a sandbox's expiration time by its ID.
func (c *Client) ExtendSandbox(ctx context.Context, sandboxId string, dto ExtendSandboxDto) error {
	c.config.Logger.Info("Extending sandbox", "sandboxId:", sandboxId, "dto:", dto)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/extend", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to extend sandbox: %s", resp.Status)
	}

	return nil
}

// ExecuteComputerUseAction executes a computer use action on the sandbox.
func (c *Client) ExecuteComputerUseAction(ctx context.Context, sandboxId string, dto ComputerUseActionDto) (*SandboxActionResponseDto, error) {
	c.config.Logger.Info("Executing computer use action", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/actions/computer-use", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to execute computer use action: %s", resp.Status)
	}

	var actionResponse SandboxActionResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&actionResponse); err != nil {
		return nil, err
	}

	return &actionResponse, nil
}

// PreviewSandbox takes a screenshot and gets the cursor position of the sandbox.
func (c *Client) PreviewSandbox(ctx context.Context, sandboxId string) (*SandboxActionResponseDto, error) {
	c.config.Logger.Info("Previewing sandbox", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/preview", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to preview sandbox: %s", resp.Status)
	}

	var preview SandboxActionResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&preview); err != nil {
		return nil, err
	}

	return &preview, nil
}

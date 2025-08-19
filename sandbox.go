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

// ListSandboxes returns a list of sandboxes for the organization.
func (c *client) ListSandboxes(ctx context.Context) ([]GetSandboxResponseDtoSandbox, error) {
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
func (c *client) CreateSandbox(ctx context.Context, dto CreateSandboxDto) (*GetSandboxResponseDto, error) {
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
func (c *client) GetSandbox(ctx context.Context, sandboxId string) (*GetSandboxResponseDto, error) {
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
func (c *client) DeleteSandbox(ctx context.Context, sandboxId string) error {
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
func (c *client) ExtendSandbox(ctx context.Context, sandboxId string, dto ExtendSandboxDto) error {
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
func (c *client) ExecuteComputerUseAction(ctx context.Context, sandboxId string, dto ComputerUseActionDto) (*SandboxActionResponseDto, error) {
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
func (c *client) PreviewSandbox(ctx context.Context, sandboxId string) (*SandboxActionResponseDto, error) {
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

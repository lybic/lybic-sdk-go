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
	"fmt"
	"net/http"
)

// ListSandboxes returns a list of sandboxes for the organization.
func (c *client) ListSandboxes(ctx context.Context) ([]CreateSandboxResponseDto, error) {
	c.config.Logger.Info("Listing sandboxes")

	url := fmt.Sprintf("/api/orgs/%s/sandboxes", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}

	var sandboxes []CreateSandboxResponseDto
	if err := tryToGetDto[[]CreateSandboxResponseDto](resp, &sandboxes); err != nil {
		return nil, err
	}

	return sandboxes, nil
}

// CreateSandbox creates a new sandbox.
func (c *client) CreateSandbox(ctx context.Context, dto CreateSandboxDto) (*CreateSandboxResponseDto, error) {
	c.config.Logger.Info("Creating sandbox", "dto:", dto)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}

	var sandbox CreateSandboxResponseDto
	if err := tryToGetDto[CreateSandboxResponseDto](resp, &sandbox); err != nil {
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

	var sandbox GetSandboxResponseDto
	if err := tryToGetDto[GetSandboxResponseDto](resp, &sandbox); err != nil {
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

	return tryToGetDto[any](resp, nil)
}

// ExtendSandbox extends a sandbox's expiration time by its ID.
func (c *client) ExtendSandbox(ctx context.Context, sandboxId string, dto ExtendSandboxDto) error {
	c.config.Logger.Info("Extending sandbox", "sandboxId:", sandboxId, "dto:", dto)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/extend", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return err
	}

	return tryToGetDto[any](resp, nil)
}

// ExecuteComputerUseAction executes a computer use action on the sandbox.
//
//	Deprecated: Use ExecuteSandboxAction instead.
func (c *client) ExecuteComputerUseAction(ctx context.Context, sandboxId string, dto ComputerUseActionDto) (*SandboxActionResponseDto, error) {
	c.config.Logger.Info("Executing computer use action", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/actions/computer-use", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}

	var actionResponse SandboxActionResponseDto
	if err := tryToGetDto[SandboxActionResponseDto](resp, &actionResponse); err != nil {
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

	var preview SandboxActionResponseDto
	if err := tryToGetDto[SandboxActionResponseDto](resp, &preview); err != nil {
		return nil, err
	}

	return &preview, nil
}
func (c *client) ExecuteSandboxAction(ctx context.Context, sandboxId string, dto ExecuteSandboxActionDto) (*SandboxActionResponseDto, error) {
	c.config.Logger.Info("Executes a computer use or mobile use action on the sandbox", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/actions/execute", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}

	var actionResponse SandboxActionResponseDto
	if err := tryToGetDto[SandboxActionResponseDto](resp, &actionResponse); err != nil {
		return nil, err
	}

	return &actionResponse, nil
}

// CopyFilesWithSandbox copies files to/from the sandbox.
func (c *client) CopyFilesWithSandbox(ctx context.Context, sandboxId string, dto SandboxFileCopyRequestDto) (*SandboxFileCopyResponseDto, error) {
	c.config.Logger.Info("Copying files with sandbox", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/file/copy", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}

	var copyResponse SandboxFileCopyResponseDto
	if err := tryToGetDto[SandboxFileCopyResponseDto](resp, &copyResponse); err != nil {
		return nil, err
	}

	return &copyResponse, nil
}

// ExecSandboxProcess executes a process inside the sandbox.
func (c *client) ExecSandboxProcess(ctx context.Context, sandboxId string, dto SandboxProcessRequestDto) (*SandboxProcessResponseDto, error) {
	c.config.Logger.Info("Executing process in sandbox", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/process", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}

	var processResponse SandboxProcessResponseDto
	if err := tryToGetDto[SandboxProcessResponseDto](resp, &processResponse); err != nil {
		return nil, err
	}

	return &processResponse, nil
}

// CreateSandboxFromImage creates a new sandbox from a machine image.
func (c *client) CreateSandboxFromImage(ctx context.Context, dto CreateSandboxFromImageDto) (*CreateSandboxFromImageResponseDto, error) {
	c.config.Logger.Info("Creating sandbox from image", "dto:", dto)
	if dto.MaxLifeSeconds <= 0 {
		c.config.Logger.Warn("maxLifeSeconds is invalid, set to 3600")
		dto.MaxLifeSeconds = 3600
	}

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/from-image", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}
	var sandbox CreateSandboxFromImageResponseDto
	if err := tryToGetDto[CreateSandboxFromImageResponseDto](resp, &sandbox); err != nil {
		return nil, err
	}

	return &sandbox, nil
}

// GetSandboxStatus returns the status of a sandbox (PENDING/RUNNING/STOPPED/ERROR).
func (c *client) GetSandboxStatus(ctx context.Context, sandboxId string) (*SandboxStatusDto, error) {
	c.config.Logger.Info("Getting sandbox status", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/status", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}

	var status SandboxStatusDto
	if err := tryToGetDto[SandboxStatusDto](resp, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

func (c *client) Restart(ctx context.Context, sandboxId string) error {
	c.config.Logger.Info("Restarting sandbox", "sandboxId:", sandboxId)

	url := fmt.Sprintf("/api/orgs/%s/sandboxes/%s/restart", c.config.OrgId, sandboxId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, nil)
	if err != nil {
		return err
	}
	return tryToGetDto[any](resp, nil)
}

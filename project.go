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

// ListProjectsResponseDto is the response DTO for listing projects.
// NOTE: This is a placeholder as the actual response is not defined in the OpenAPI spec.

// ListProjects returns a list of projects for the organization.
func (c *client) ListProjects(ctx context.Context) ([]SingleProjectResponseDto, error) {
	url := fmt.Sprintf("/api/orgs/%s/projects", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to list projects: %s", resp.Status)
	}

	var projects []SingleProjectResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// CreateProject creates a new project.
func (c *client) CreateProject(ctx context.Context, dto CreateProjectDto) (*SingleProjectResponseDto, error) {
	url := fmt.Sprintf("/api/orgs/%s/projects", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create project: %s", resp.Status)
	}

	var project SingleProjectResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return nil, err
	}

	return &project, nil
}

// DeleteProject deletes a project by its ID.
func (c *client) DeleteProject(ctx context.Context, projectId string) error {
	url := fmt.Sprintf("/api/orgs/%s/projects/%s", c.config.OrgId, projectId)
	resp, err := c.request(ctx, http.MethodDelete, url, nil, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete project: %s", resp.Status)
	}

	return nil
}

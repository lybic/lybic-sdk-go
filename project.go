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
func (c *Client) ListProjects(ctx context.Context) ([]SingleProjectResponseDto, error) {
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
func (c *Client) CreateProject(ctx context.Context, dto CreateProjectDto) (*SingleProjectResponseDto, error) {
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
func (c *Client) DeleteProject(ctx context.Context, projectId string) error {
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

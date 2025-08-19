package lybic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetStats returns the stats of the organization.
func (c *Client) GetStats(ctx context.Context) (*StatsResponseDto, error) {
	c.config.Logger.Info("Getting organization stats")

	url := fmt.Sprintf("/api/orgs/%s/stats", c.config.OrgId)
	resp, err := c.request(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.config.Logger.Errorf("failed to get stats: %s", resp.Status)
		return nil, fmt.Errorf("failed to get stats: %s", resp.Status)
	}

	var stats StatsResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		c.config.Logger.Errorf("failed to decode stats response: %v", err)
		return nil, err
	}

	return &stats, nil
}

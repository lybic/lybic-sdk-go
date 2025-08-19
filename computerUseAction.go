package lybic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ParseComputerUse parses the output text of a computer use model and returns the parsed actions.
func (c *Client) ParseComputerUse(ctx context.Context, dto ComputerUseParseRequestDto) (*ComputerUseActionResponseDto, error) {
	url := "/api/computer-use/parse"
	c.config.Logger.Info("Sending request to parse computer use action", "url:", url, "dto:", dto)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.config.Logger.Errorf("failed to parse computer use: %s", resp.Status)
		// Log the response body for debugging purposes
		body, _ := io.ReadAll(resp.Body)
		c.config.Logger.Errorf("Response body: %s", body)
		return nil, fmt.Errorf("failed to parse computer use: %s", resp.Status)
	}

	var actions ComputerUseActionResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&actions); err != nil {
		c.config.Logger.Errorf("failed to decode response body: %v", err)
		return nil, err
	}

	return &actions, nil
}

package solus

import (
	"context"
	"encoding/json"
	"fmt"
)

type LocationCreateRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Icon             string `json:"icon"`
	IsDefault        bool   `json:"is_default"`
	IsVisible        bool   `json:"is_visible"`
	ComputeResources []int  `json:"compute_resources"`
}

type Location struct {
	Id               int               `json:"id"`
	Name             string            `json:"name"`
	Icon             string            `json:"icon"`
	Description      string            `json:"description"`
	IsDefault        bool              `json:"is_default"`
	IsVisible        bool              `json:"is_visible"`
	ComputeResources []ComputeResource `json:"compute_resources"`
}

type LocationCreateResponse struct {
	Data Location `json:"data"`
}

func (c *Client) LocationCreate(ctx context.Context, data LocationCreateRequest) (Location, error) {
	body, code, err := c.request(ctx, "POST", "locations", data)
	if err != nil {
		return Location{}, err
	}

	if code != 201 {
		return Location{}, fmt.Errorf("HTTP %d: %s", code, body)
	}

	var resp LocationCreateResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return Location{}, fmt.Errorf("failed to decode '%s': %s", body, err)
	}

	return resp.Data, nil
}

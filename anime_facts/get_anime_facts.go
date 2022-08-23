package anime_facts

import (
	"encoding/json"
	"fmt"
	"io"
)

type animeFacts struct {
	Success bool   `json:"success"`
	Total   uint64 `json:"total_facts"`
	Image   uint64 `json:"anime_img"`
	Data    []Fact `json:"data"`
}

func (c *Client) GetAnimeFacts(animeName string) ([]Fact, error) {
	url := fmt.Sprintf("%s/%s", c.baseUrl, animeName)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r animeFacts
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	if !r.Success {
		return nil, ErrExternalService
	}

	return r.Data, nil
}

package anime_facts

import (
	"encoding/json"
	"fmt"
	"io"
)

type animeSpecificFact struct {
	Success bool `json:"success"`
	Data    Fact `json:"data"`
}

func (c *Client) GetSpecificFact(animeName string, factId uint64) (*Fact, error) {
	url := fmt.Sprintf("%s/%s/%d", c.baseUrl, animeName, factId)
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r animeSpecificFact
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	if !r.Success {
		return nil, ErrExternalService
	}

	return &r.Data, nil
}

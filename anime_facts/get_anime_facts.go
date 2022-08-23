package anime_facts

import (
	"fmt"
)

type animeFacts struct {
	Success bool   `json:"success"`
	Total   uint64 `json:"total_facts"`
	Image   uint64 `json:"anime_img"`
	Data    []Fact `json:"data"`
}

func (c *Client) GetAnimeFacts(animeName string) ([]Fact, error) {
	url := fmt.Sprintf("%s/%s", c.baseUrl, animeName)

	var r animeFacts
	if err := c.GetRequest(url, &r); err != nil {
		return nil, err
	}

	if !r.Success {
		return nil, ErrExternalService
	}

	return r.Data, nil
}

package anime_facts

import (
	"fmt"
)

type animeSpecificFact struct {
	Success bool `json:"success"`
	Data    Fact `json:"data"`
}

func (c *Client) GetSpecificFact(animeName string, factId uint64) (*Fact, error) {
	url := fmt.Sprintf("%s/%s/%d", c.baseUrl, animeName, factId)

	var r animeSpecificFact
	if err := c.GetRequest(url, &r); err != nil {
		return nil, err
	}

	if !r.Success {
		return nil, ErrExternalService
	}

	return &r.Data, nil
}

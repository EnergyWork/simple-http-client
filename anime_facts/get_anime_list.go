package anime_facts

import (
	"encoding/json"
	"io"
)

type animeList struct {
	Success bool    `json:"success"`
	Data    []Anime `json:"data"`
}

func (c *Client) GetAnimeList() ([]Anime, error) {
	url := BaseApiURL
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r animeList
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	if !r.Success {
		return nil, ErrExternalService
	}

	return r.Data, nil
}

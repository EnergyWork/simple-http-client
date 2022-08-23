package anime_facts

type animeList struct {
	Success bool    `json:"success"`
	Data    []Anime `json:"data"`
}

func (c *Client) GetAnimeList() ([]Anime, error) {
	url := BaseApiURL

	var r animeList
	if err := c.GetRequest(url, &r); err != nil {
		return nil, err
	}

	if !r.Success {
		return nil, ErrExternalService
	}

	return r.Data, nil
}

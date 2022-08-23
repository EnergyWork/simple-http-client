package anime_facts

import "fmt"

type Anime struct {
	ID    uint64 `json:"anime_id"`
	Name  string `json:"anime_name"`
	Image string `json:"anime_img"` // url to external resource
}

func (a *Anime) String() string {
	return fmt.Sprintf("ID: %d | Name: %s | Image: %s", a.ID, a.Name, a.Image)
}

type Fact struct {
	ID   uint64 `json:"fact_id"`
	Fact string `json:"fact"`
}

func (f *Fact) String() string {
	return fmt.Sprintf("ID: %d | Fact: %s", f.ID, f.Fact)
}

package anime_facts

import "errors"

const (
	BaseApiURL = "https://anime-facts-rest-api.herokuapp.com/api/v1"
)

var (
	ErrExternalService = errors.New("the external service responded with an error")
)

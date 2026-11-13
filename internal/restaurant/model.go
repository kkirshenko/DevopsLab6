package restaurant

import "errors"

var (
	ErrNotFound = errors.New("restaurant not found")
	ErrInvalid  = errors.New("invalid restaurant")
)

type Restaurant struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Rating  float32 `json:"rating"`
	Open    bool    `json:"open"`
}

package restaurant

import "context"

type Repository interface {
	Create(context.Context, Restaurant) (string, error)
	GetByID(context.Context, string) (Restaurant, error)
	ChangeRating(context.Context, string, float32) error
}

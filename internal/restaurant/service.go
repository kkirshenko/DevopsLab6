package restaurant

import (
	"context"
	"fmt"
)

type Service interface {
	Create(context.Context, Restaurant) (string, error)
	GetByID(context.Context, string) (Restaurant, error)
	ChangeRating(context.Context, string, float32) error
}

type service struct { repo Repository }

func NewService(repo Repository) Service { return &service{repo: repo} }

func (s *service) Create(ctx context.Context, item Restaurant) (string, error) {
	if item.ID == "" || item.Name == "" || item.Address == "" || item.Rating < 0 || item.Rating > 5 {
		return "", fmt.Errorf("%w: id, name, address are required and rating must be 0..5", ErrInvalid)
	}
	return s.repo.Create(ctx, item)
}
func (s *service) GetByID(ctx context.Context, id string) (Restaurant, error) {
	if id == "" { return Restaurant{}, ErrInvalid }
	return s.repo.GetByID(ctx, id)
}
func (s *service) ChangeRating(ctx context.Context, id string, rating float32) error {
	if id == "" || rating < 0 || rating > 5 { return ErrInvalid }
	return s.repo.ChangeRating(ctx, id, rating)
}

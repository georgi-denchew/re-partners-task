package domain

import (
	"context"

	"repartnerstask.com/m/internal/domain/types"
)

//go:generate go tool mockgen -destination=./mocks/types.go -typed -package mocks . Repository
type Repository interface {
	StoreOrder(ctx context.Context, order *types.Order) error
	ListOrdersDesc(ctx context.Context, count int) ([]*types.Order, error)
}

type Application struct {
	packSizes []int
	repo      Repository
}

func NewApplication(availablePackSizes []int, repo Repository) (Application, error) {
	if len(availablePackSizes) == 0 {
		return Application{}, NewNoPacksError()
	}
	app := Application{
		packSizes: availablePackSizes,
		repo:      repo,
	}

	return app, nil
}

package handlers

import (
	"context"

	"repartnerstask.com/m/internal/domain/types"
)

//go:generate go tool mockgen -destination=./mocks/types.go -typed -package mocks . Application
type Application interface {
	CreateOrder(ctx context.Context, orderItemsCount int) (packsSizeCount map[int]int, err error)
	GetOrders(ctx context.Context, ordersCount int) ([]*types.Order, error)
	ReplacePacks(packs []int)
}

type Server struct {
	app Application
}

func NewServer(app Application) *Server {
	return &Server{
		app: app,
	}
}

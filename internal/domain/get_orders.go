package domain

import (
	"context"

	"repartnerstask.com/m/internal/domain/types"
)

func (a *Application) GetOrders(ctx context.Context, ordersCount int) ([]*types.Order, error) {
	return a.repo.ListOrdersDesc(ctx, ordersCount)
}

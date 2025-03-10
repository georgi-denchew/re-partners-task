package repository

import (
	"context"
	"sync"

	"repartnerstask.com/m/internal/domain/types"
)

// InMemoryStorage is only for demonstration purposes. Should be replaced with a DB.
type InMemoryStorage struct {
	// using a Mutex to make the in-memory-storage thread-safe
	m      sync.Mutex
	orders []*types.Order
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		orders: make([]*types.Order, 0),
	}
}

// CreateOrder stores a order in the in-memory storage. Returns nil if successful, otherwise returns an error.
func (i *InMemoryStorage) StoreOrder(ctx context.Context, order *types.Order) error {
	// ideally this function should receive a 'repository.Order' parameter which can have repository-related metadata
	i.m.Lock()
	defer i.m.Unlock()

	i.orders = append(i.orders, order)

	return nil
}

// ListOrdersDesc returns the last stored orders. Number of orders returned is specified by the count param.
// Returns a slice if successful, otherwise returns an error.
func (i *InMemoryStorage) ListOrdersDesc(ctx context.Context, count int) ([]*types.Order, error) {
	// ideally this function should receive a repository.Order parameter which can have repository-related metadata
	i.m.Lock()
	defer i.m.Unlock()

	result := make([]*types.Order, 0)
	index := len(i.orders) - 1

	for index >= 0 && len(result) < count {
		result = append(result, i.orders[index])
		index--
	}

	return result, nil
}

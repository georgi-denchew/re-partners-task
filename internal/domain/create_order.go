package domain

import (
	"context"
	"math"
	"time"

	"repartnerstask.com/m/internal/domain/types"
)

// CreateOrder calculates items and packs needed to fulfill the order and interacts with a Repository to store it.
// If successful, returns a map[int]int where the key is items-per-pack and value is packs-count.
// Returns an error on failure.
func (a *Application) CreateOrder(ctx context.Context, orderItemsCount int) (map[int]int, error) {
	biggestPack := 0

	for _, packSize := range a.packSizes {
		if biggestPack < packSize {
			biggestPack = packSize
		}
	}

	itemsCountToCheck := orderItemsCount + biggestPack

	ordersPerItemCount := make(map[int]*types.Order, itemsCountToCheck)
	ordersPerItemCount[0] = &types.Order{
		ItemsShipped: 0,
		PacksShipped: 0,
	}

	for i := 1; i <= itemsCountToCheck; i++ {
		ordersPerItemCount[i] = &types.Order{ItemsShipped: math.MaxInt32, PacksShipped: math.MaxInt32}
		for _, packSize := range a.packSizes {
			if i >= packSize && ordersPerItemCount[i-packSize].ItemsShipped != math.MaxInt32 {
				if ordersPerItemCount[i].ItemsShipped > ordersPerItemCount[i-packSize].ItemsShipped ||
					(ordersPerItemCount[i].ItemsShipped == ordersPerItemCount[i-packSize].ItemsShipped &&
						ordersPerItemCount[i].PacksShipped > ordersPerItemCount[i-packSize].PacksShipped) {

					ordersPerItemCount[i].ItemsShipped = ordersPerItemCount[i-packSize].ItemsShipped + packSize
					ordersPerItemCount[i].PacksShipped = ordersPerItemCount[i-packSize].PacksShipped + 1

					ordersPerItemCount[i].ItemPacks = make(map[int]int)
					for itemsPerPack, packCount := range ordersPerItemCount[i-packSize].ItemPacks {
						ordersPerItemCount[i].ItemPacks[itemsPerPack] = packCount
					}

					ordersPerItemCount[i].ItemPacks[packSize]++
				}
			}
		}
	}

	var chosenOrder *types.Order
	for minShippedItems := orderItemsCount; minShippedItems <= itemsCountToCheck; minShippedItems++ {
		if ordersPerItemCount[minShippedItems].ItemsShipped != math.MaxInt32 {
			chosenOrder = ordersPerItemCount[minShippedItems]
			break
		}
	}

	if chosenOrder == nil {
		return nil, &CannotCalculateOrderItemsError{}
	}

	chosenOrder.CreatedAt = time.Now().UTC()

	if err := a.repo.StoreOrder(ctx, chosenOrder); err != nil {
		return nil, err
	}

	return chosenOrder.ItemPacks, nil
}

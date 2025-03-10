package types

import "time"

type Order struct {
	ItemsShipped int
	PacksShipped int
	ItemPacks    map[int]int
	CreatedAt    time.Time
}

package handlers

import "time"

type CreateOrderRequest struct {
	ItemsCount int `json:"items_count"`
}

type OrderResponse struct {
	CreatedAt *time.Time  `json:"created_at,omitempty"`
	Packs     []*ItemPack `json:"item_packs"`
}

type ItemPack struct {
	Items int `json:"items"`
	Packs int `json:"packs"`
}

type ReplacePacksRequest struct {
	Packs []int `json:"packs"`
}

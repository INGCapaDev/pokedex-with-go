package config

import (
	"github.com/ingcapadev/pokedex-with-go/internal/items"
)

type InventoryItem interface {
	GetBaseInfo() items.ItemBaseInfo
	GetSellPrice() float64
	GetCanBeSold() bool
	GetQuantity() int
	SetQuantity(int)
}

type Inventory struct {
	Items       map[string]InventoryItem `json:"items"`
	Balance     float64                  `json:"balance"`
	MaxCapacity int                      `json:"max_capacity"`
}

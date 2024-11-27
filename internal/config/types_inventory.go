package config

type ItemBaseInfo struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type InventoryItemInfo struct {
	ItemBaseInfo
	Quantity     int     `json:"quantity"`
	IsConsumable bool    `json:"is_consumable"`
	CanBeSold    bool    `json:"can_be_sold"`
	SellPrice    float64 `json:"sell_price"`
}

type Inventory struct {
	Items       map[string]InventoryItemInfo `json:"items"`
	Balance     float64                      `json:"balance"`
	MaxCapacity int                          `json:"max_capacity"`
}

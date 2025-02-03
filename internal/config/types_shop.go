package config

type ShopItem struct {
	Item  InventoryItem `json:"item"`
	Price float64       `json:"price"`
}

type Shop struct {
	Items         map[string]ShopItem `json:"items"`
	DevaluingRate float64             `json:"devaluing_rate"`
}

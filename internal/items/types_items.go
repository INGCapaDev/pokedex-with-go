package items

type ItemBaseInfo struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsConsumable bool   `json:"is_consumable"`
}

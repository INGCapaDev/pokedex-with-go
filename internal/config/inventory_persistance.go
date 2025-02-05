package config

import (
	"encoding/json"
	"fmt"

	"github.com/ingcapadev/pokedex-with-go/internal/items"
)

type TempInventory struct {
	Items       map[string]json.RawMessage `json:"items"`
	Balance     float64                    `json:"balance"`
	MaxCapacity int                        `json:"max_capacity"`
}

func (inv *Inventory) UnmarshalJSON(data []byte) error {
	var temp TempInventory
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	inv.Items = make(map[string]InventoryItem)
	inv.Balance = temp.Balance
	inv.MaxCapacity = temp.MaxCapacity

	for name, itemData := range temp.Items {
		var item InventoryItem
		if err := json.Unmarshal(itemData, &item); err == nil {
			inv.Items[name] = item
			continue
		}

		var pokeball items.Pokeball
		if err := json.Unmarshal(itemData, &pokeball); err == nil {
			inv.Items[name] = &pokeball
			continue
		}

		return fmt.Errorf("failed to unmarshal item %s: %v", name, "unknown item type or invalid data")
	}

	return nil
}

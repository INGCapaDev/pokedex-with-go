package config

import (
	"errors"
	"fmt"
)

func (inv *Inventory) canAddToInventory() error {
	if len(inv.Items) >= inv.MaxCapacity {
		return errors.New("Inventory is full, you cannot add more items, try selling or using some items; or increase the inventory capacity")
	}
	return nil
}

func (inv *Inventory) haveTheRequiredBalance(amount float64) bool {
	return inv.Balance >= amount
}

func (inv *Inventory) hasTheRequiredSpace(quantity int) bool {
	availableSpace := inv.GetAvailableInventorySpace()
	return availableSpace >= quantity
}

func (inv *Inventory) existsInInventory(itemName string) bool {
	_, exists := inv.Items[itemName]
	return exists
}

func (inv *Inventory) canSellItem(itemName string, quantity int) error {
	if !inv.existsInInventory(itemName) {
		return fmt.Errorf("You don't have %s in your inventory", itemName)
	}

	if !inv.Items[itemName].GetCanBeSold() {
		return fmt.Errorf("%s cannot be sold", itemName)
	}

	if inv.Items[itemName].GetQuantity() <= quantity {
		return fmt.Errorf("You only have %d %s in your inventory", inv.Items[itemName].GetQuantity(), itemName)
	}

	return nil
}

func (inv *Inventory) canUseItem(itemName string) error {
	if !inv.existsInInventory(itemName) {
		return fmt.Errorf("You don't have %s in your inventory", itemName)
	}

	if !inv.Items[itemName].GetBaseInfo().IsConsumable {
		return fmt.Errorf("%s is not a consumable item", itemName)
	}

	if inv.Items[itemName].GetQuantity() <= 0 {
		return fmt.Errorf("You don't have any %s in your inventory", itemName)
	}

	return nil
}

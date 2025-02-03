package config

func (inv *Inventory) handleChangeInventoryItemQuantity(itemName string, quantity int) {
	if quantity <= 0 {
		delete(inv.Items, itemName)
	}
	newItem := inv.Items[itemName]
	newItem.SetQuantity(quantity)
	inv.Items[itemName] = newItem
}

func (inv *Inventory) handleAddItem(item InventoryItem) {
	if _, exists := inv.Items[item.GetBaseInfo().Name]; exists {
		newQuantity := inv.Items[item.GetBaseInfo().Name].GetQuantity() + item.GetQuantity()
		item.SetQuantity(newQuantity)
	}
	inv.Items[item.GetBaseInfo().Name] = item
}

func (inv *Inventory) GetAvailableInventorySpace() int {
	var usedSpace int
	for _, item := range inv.Items {
		usedSpace += item.GetQuantity()
	}
	return inv.MaxCapacity - usedSpace
}

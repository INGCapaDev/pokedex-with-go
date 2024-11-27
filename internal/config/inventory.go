package config

import (
	"encoding/json"
	"errors"
)

func loadInventoryFromDisk(cfg *TConfig) error {
	fileData, err, notFound := readDataFromDisk(INVENTORY_FILE)
	if err != nil {
		if notFound {
			return nil
		}
		return err
	}

	return json.Unmarshal(fileData, &cfg.Inventory)
}

func (cfg *TConfig) AddToInventory(item InventoryItemInfo) error {
	if _, exists := cfg.Inventory.Items[item.Name]; exists {
		newQuantity := cfg.Inventory.Items[item.Name].Quantity + item.Quantity
		item.Quantity = newQuantity
	}

	cfg.Inventory.Items[item.Name] = item

	err := writeDataToDisk(cfg.Inventory, INVENTORY_FILE)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *TConfig) UseItem(itemName string) error {
	notFound := errors.New("Item not found in inventory")
	if _, exists := cfg.Inventory.Items[itemName]; !exists {
		return notFound
	}

	if !cfg.Inventory.Items[itemName].IsConsumable {
		return errors.New("Item is not consumable")
	}

	if cfg.Inventory.Items[itemName].Quantity <= 0 {
		delete(cfg.Inventory.Items, itemName)
		return notFound
	}

	newItem := cfg.Inventory.Items[itemName]
	newItem.Quantity--

	if newItem.Quantity <= 0 {
		delete(cfg.Inventory.Items, itemName)
	}

	cfg.Inventory.Items[itemName] = newItem
	if newItem.Quantity <= 0 {
		delete(cfg.Inventory.Items, itemName)
	}

	err := writeDataToDisk(cfg.Inventory, INVENTORY_FILE)
	if err != nil {
		return err
	}

	return nil
}

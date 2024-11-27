package config

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	INITIAL_BALANCE  = 500
	INITIAL_CAPACITY = 10
)

func loadInventoryFromDisk(cfg *TConfig) error {
	fileData, err, notFound := readDataFromDisk(INVENTORY_FILE)
	if err != nil {
		if notFound {
			fmt.Println("PokeCLI: Inventory file not found. Initializing inventory...")
			err = cfg.initInventory()
			if err != nil {
				return err
			}
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

func (cfg *TConfig) initInventory() error {
	fmt.Println("PokeCLI: Initializing inventory...")
	pokedex := getPokedexInfo()
	pokeball := getPokeballInfo()

	err := cfg.AddToInventory(InventoryItemInfo{
		ItemBaseInfo: pokedex,
		Quantity:     1,
		IsConsumable: false,
		CanBeSold:    pokedex.Price > 0,
		SellPrice:    pokedex.Price * 0.7,
	})
	if err != nil {
		return err
	}

	err = cfg.AddToInventory(InventoryItemInfo{
		ItemBaseInfo: pokeball,
		Quantity:     10,
		IsConsumable: true,
		CanBeSold:    pokeball.Price > 0,
		SellPrice:    pokeball.Price * 0.7,
	})
	if err != nil {
		return err
	}

	err = cfg.setInitialValues()
	if err != nil {
		return err
	}

	fmt.Println("PokeCLI: Inventory initialized successfully")
	return nil
}

func (cfg *TConfig) setInitialValues() error {
	cfg.Inventory.Balance = INITIAL_BALANCE
	cfg.Inventory.MaxCapacity = INITIAL_CAPACITY
	err := writeDataToDisk(cfg.Inventory, INVENTORY_FILE)
	if err != nil {
		return err
	}
	return nil
}

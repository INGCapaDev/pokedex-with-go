package config

import (
	"encoding/json"
	"fmt"

	"github.com/ingcapadev/pokedex-with-go/internal/items"
)

const (
	INITIAL_BALANCE  = 500
	INITIAL_CAPACITY = 20
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

func (cfg *TConfig) UseItem(itemName string) error {

	err := cfg.Inventory.canUseItem(itemName)
	if err != nil {
		return err
	}

	item := cfg.Inventory.Items[itemName]
	cfg.Inventory.handleChangeInventoryItemQuantity(itemName, item.GetQuantity()-1)

	err = writeDataToDisk(cfg.Inventory, INVENTORY_FILE)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *TConfig) SellItem(itemName string, quantity int) error {

	err := cfg.Inventory.canSellItem(itemName, quantity)
	if err != nil {
		return err
	}

	newQuantity := cfg.Inventory.Items[itemName].GetQuantity() - quantity
	cfg.Inventory.handleChangeInventoryItemQuantity(itemName, newQuantity)
	cfg.Inventory.Balance += float64(quantity) * cfg.Inventory.Items[itemName].GetSellPrice()

	err = writeDataToDisk(cfg.Inventory, INVENTORY_FILE)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *TConfig) BuyItem(name string, quantity int) error {

	item, err := cfg.Shop.Purchase(name, quantity)
	if err != nil {
		return err
	}

	priceToPay := float64(item.Item.GetQuantity()) * item.Price
	if !cfg.Inventory.haveTheRequiredBalance(priceToPay) {
		required := priceToPay - cfg.Inventory.Balance
		return fmt.Errorf("insufficient funds. You need $%.2f more, to buy %d %s", required, item.Item.GetQuantity(), item.Item.GetBaseInfo().Name)

	}
	if !cfg.Inventory.hasTheRequiredSpace(item.Item.GetQuantity()) {
		required := item.Item.GetQuantity() - cfg.Inventory.GetAvailableInventorySpace()
		return fmt.Errorf("insufficient space. You only	have space for %d more items", required)
	}

	cfg.Inventory.Balance -= priceToPay
	cfg.Inventory.handleAddItem(item.Item)

	err = writeDataToDisk(cfg.Inventory, INVENTORY_FILE)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *TConfig) initInventory() error {
	fmt.Println("PokeCLI: Initializing inventory...")
	pokedex := items.CreatePokedex()
	pokeballs := items.CreatePokeball(10, true, 0, items.ITEM_POKEBALL)

	cfg.Inventory.handleAddItem(pokeballs)
	cfg.Inventory.handleAddItem(pokedex)
	cfg.Inventory.Balance = INITIAL_BALANCE
	cfg.Inventory.MaxCapacity = INITIAL_CAPACITY

	err := writeDataToDisk(cfg.Inventory, INVENTORY_FILE)
	if err != nil {
		return err
	}

	fmt.Println("PokeCLI: Inventory initialized successfully")
	return nil
}

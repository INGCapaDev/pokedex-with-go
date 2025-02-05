package main

import (
	"fmt"

	"github.com/ingcapadev/pokedex-with-go/internal/config"
)

func printHelpShop(_ *config.TConfig, _ ...string) error {
	fmt.Printf("\nUsage: %s <command>", SHOP_CMD)
	for _, cmd := range getShopCommands() {
		fmt.Printf("\n%s: %s", cmd.name, cmd.description)
	}
	fmt.Printf("\n\n")
	return nil
}

func cmdShop(cfg *config.TConfig, args ...string) error {
	if len(args) == 0 {
		return printHelpShop(cfg)
	}

	command, exists := getShopCommands()[args[0]]
	if !exists {
		return fmt.Errorf("command not found. Type '%s %s' for help", HELP_CMD, SHOP_CMD)
	}
	err := command.callback(cfg, args[1:]...)
	if err != nil {
		return fmt.Errorf("error executing %s command: %v", command.name, err)
	}
	return nil

}

func cmdShopList(cfg *config.TConfig, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("Invalid number of arguments. Usage: %s %s", SHOP_CMD, SHOP_CMD_LIST)
	}

	cfg.Shop.PrintShop()
	return nil
}

func askForConfirmation() bool {
	fmt.Printf("\nAre you sure you want to buy the item? (Y/n): ")
	var confirmation string
	fmt.Scanln(&confirmation)
	return confirmation == "y" || confirmation == "Y" || confirmation == ""
}

func askForQuantity() (int, error) {
	fmt.Printf("\nHow many items do you want to buy? : ")
	var quantity int
	_, err := fmt.Scanln(&quantity)
	if err != nil {
		return 0, err
	}
	if quantity <= 0 {
		return 0, fmt.Errorf("Quantity must be greater than 0")
	}
	return quantity, nil
}

func cmdShopBuy(cfg *config.TConfig, args ...string) error {
	if len(args) == 0 {
		fmt.Printf("\nUsage: %s %s <item-name>", SHOP_CMD, SHOP_CMD_BUY)
		return fmt.Errorf("You must provide an item to buy")
	}

	if len(args) != 1 {
		fmt.Printf("\nUsage: %s %s <item-name>", SHOP_CMD, SHOP_CMD_BUY)
		return fmt.Errorf("You must provide an item to buy")
	}

	//ask for confirmation to buy the item
	confirmed := askForConfirmation()
	if !confirmed {
		fmt.Printf("\nOperation cancelled.\n\n")
		return nil
	}

	//ask for quantity
	quantity, err := askForQuantity()
	if err != nil {
		return fmt.Errorf("Invalid quantity. %v, try again with a valid integer", err)
	}

	fmt.Printf("\nBuying item: %s... \n", args[0])
	err = cfg.BuyItem(args[0], quantity)
	if err != nil {
		return fmt.Errorf("Error buying item: %v", err)
	}
	fmt.Printf("\n%s bought successfully!\n\n", args[0])

	return nil
}

func getShopCommands() map[string]cliCommand {
	return map[string]cliCommand{
		SHOP_CMD_BUY: {
			name:        SHOP_CMD_BUY,
			description: "Buy an item from the shop",
			callback:    cmdShopBuy,
		},
		SHOP_CMD_LIST: {
			name:        SHOP_CMD_LIST,
			description: "Sell an item from the inventory",
			callback:    cmdShopList,
		},
	}
}

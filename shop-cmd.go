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
	fmt.Print("Shop list\n")
	return nil
}

func cmdShopBuy(cfg *config.TConfig, args ...string) error {

	if len(args) == 0 {

		fmt.Printf("\nUsage: %s %s <item-name>", SHOP_CMD, SHOP_CMD_BUY)
		return fmt.Errorf("You must provide an item to buy")
	}
	fmt.Print("Shop buy\n")
	fmt.Print("Buying item: ", args[0])
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

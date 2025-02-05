package config

import (
	"fmt"
	"strings"

	"github.com/ingcapadev/pokedex-with-go/internal/items"
)

const (
	INITIAL_DEVALUING_RATE = 0.3
	POKEBALL_PRICE         = 25
)

func NewShop() *Shop {
	shop := &Shop{}
	shop.load()
	return shop
}

func (s *Shop) load() {

	s.DevaluingRate = INITIAL_DEVALUING_RATE
	s.Items = make(map[string]ShopItem)

	s.Items[strings.ToLower(string(items.ITEM_POKEBALL))] = ShopItem{
		Item:  items.CreatePokeball(1, true, POKEBALL_PRICE*(1-s.DevaluingRate), items.ITEM_POKEBALL),
		Price: 25,
	}

	s.Items[strings.ToLower(string(items.ITEM_GREATBALL))] = ShopItem{
		Item:  items.CreatePokeball(1, true, POKEBALL_PRICE*(1-s.DevaluingRate), items.ITEM_GREATBALL),
		Price: 50,
	}

	s.Items[strings.ToLower(string(items.ITEM_ULTRABALL))] = ShopItem{
		Item:  items.CreatePokeball(1, true, POKEBALL_PRICE*(1-s.DevaluingRate), items.ITEM_ULTRABALL),
		Price: 100,
	}

	s.Items[strings.ToLower(string(items.ITEM_MASTERBALL))] = ShopItem{
		Item:  items.CreatePokeball(1, true, POKEBALL_PRICE*(1-s.DevaluingRate), items.ITEM_MASTERBALL),
		Price: 500,
	}

}

func (s *Shop) Purchase(name string, quantity int) (ShopItem, error) {
	item, exists := s.Items[name]
	if !exists {
		return ShopItem{}, fmt.Errorf("item %s does not exist in the shop", name)
	}

	item.Item.SetQuantity(quantity)
	return item, nil
}

func (s *Shop) PrintShop() {
	fmt.Print("\nShop items:\n")
	position := 1
	for _, item := range s.Items {
		fmt.Printf("%d. %s - $%.2f\n", position, item.Item.GetBaseInfo().Name, item.Price)
		position++
	}
	fmt.Println()
}

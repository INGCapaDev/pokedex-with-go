package items

import (
	"math"
	"math/rand"
)

const (
	CATCH_RATE_POKEBALL   = 0.3
	CATCH_RATE_GREATBALL  = 0.45
	CATCH_RATE_ULTRABALL  = 0.6
	CATCH_RATE_MASTERBALL = 1.0
)

type Pokeball struct {
	BaseItem
	CatchRate float64 `json:"catchRate"`
}

func (p *Pokeball) TryToCatch(pokeBaseExp int) bool {

	captureChance := (p.CatchRate * 10) / math.Sqrt(float64(pokeBaseExp+1))

	randomRoll := rand.Float64()
	return randomRoll <= captureChance || p.Name == string(ITEM_MASTERBALL)
}

// Inheritance
func (p *Pokeball) GetBaseInfo() ItemBaseInfo {
	return p.BaseItem.GetBaseInfo()
}

func (p *Pokeball) GetQuantity() int {
	return p.BaseItem.GetQuantity()
}

func (p *Pokeball) SetQuantity(quantity int) {
	p.BaseItem.SetQuantity(quantity)
}

func (p *Pokeball) GetCanBeSold() bool {
	return p.BaseItem.GetCanBeSold()
}

func (p *Pokeball) GetSellPrice() float64 {
	return p.BaseItem.GetSellPrice()
}

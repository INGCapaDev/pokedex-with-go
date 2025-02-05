package items

type PokeballType string

const (
	ITEM_POKEBALL   PokeballType = "Pokeball"
	ITEM_GREATBALL  PokeballType = "Greatball"
	ITEM_MASTERBALL PokeballType = "Masterball"
	ITEM_ULTRABALL  PokeballType = "Ultraball"
	ITEM_POKEDEX                 = "Pokedex"
)

func GetPokeballInfo(pokeball PokeballType) ItemBaseInfo {
	var name string

	switch pokeball {
	case ITEM_POKEBALL:
		name = string(pokeball)
	case ITEM_GREATBALL:
		name = string(pokeball)
	case ITEM_MASTERBALL:
		name = string(pokeball)
	case ITEM_ULTRABALL:
		name = string(pokeball)
	default:
		name = string(ITEM_POKEBALL)
	}

	return ItemBaseInfo{
		Name:         name,
		IsConsumable: true,
		Description:  "A device for catching wild Pokémon. It is thrown like a ball at a Pokémon, comfortably encapsulating its target.",
	}
}

func CreatePokeball(quantity int, canBeSold bool, sellPrice float64, pokeball PokeballType) *Pokeball {
	var catchRate float64
	switch pokeball {
	case ITEM_POKEBALL:
		catchRate = CATCH_RATE_POKEBALL
	case ITEM_GREATBALL:
		catchRate = CATCH_RATE_GREATBALL
	case ITEM_ULTRABALL:
		catchRate = CATCH_RATE_ULTRABALL
	case ITEM_MASTERBALL:
		catchRate = CATCH_RATE_MASTERBALL
	default:
		catchRate = 0.3
	}
	return &Pokeball{
		BaseItem: BaseItem{
			ItemBaseInfo: GetPokeballInfo(pokeball),
			Quantity:     quantity,
			SellPrice:    sellPrice,
			CanBeSold:    canBeSold,
		},
		CatchRate: catchRate,
	}
}

func GetPokedexInfo() ItemBaseInfo {
	return ItemBaseInfo{
		Name:         ITEM_POKEDEX,
		IsConsumable: false,
		Description:  "A high-tech device that records the Pokémon you've seen and caught. It's a high-tech encyclopedia!",
	}
}

func CreatePokedex() *BaseItem {
	return &BaseItem{
		ItemBaseInfo: GetPokedexInfo(),
		Quantity:     1,
		SellPrice:    0,
		CanBeSold:    false,
	}
}

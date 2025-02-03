package items

type PokeballType string

const (
	ITEM_POKEBALL   PokeballType = "Pokeball"
	ITEM_GREATBALL  PokeballType = "Greatball"
	ITEM_MASTERBALL PokeballType = "Masterball"
	ITEM_ULTRABALL  PokeballType = "Ultraball"
	ITEM_POKEDEX                 = "Pokedex"
)

func GetPokeballInfo() ItemBaseInfo {
	return ItemBaseInfo{
		Name:         string(ITEM_POKEBALL),
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
			ItemBaseInfo: GetPokeballInfo(),
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

func GetMasterballInfo() ItemBaseInfo {
	return ItemBaseInfo{
		Name:         string(ITEM_MASTERBALL),
		IsConsumable: true,
		Description:  "The best Poké Ball with the ultimate level of performance. With it, you will catch any wild Pokémon without fail.",
	}
}

func GetGreatballInfo() ItemBaseInfo {
	return ItemBaseInfo{
		Name:         string(ITEM_GREATBALL),
		IsConsumable: true,
		Description:  "A good, high-performance Poké Ball that provides a higher Pokémon catch rate than a standard Poké Ball.",
	}
}

func GetUltraballInfo() ItemBaseInfo {
	return ItemBaseInfo{
		Name:         string(ITEM_ULTRABALL),
		IsConsumable: true,
		Description:  "An ultra-high-performance Poké Ball that provides a higher Pokémon catch rate than a Great Ball.",
	}
}

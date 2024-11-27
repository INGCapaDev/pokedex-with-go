package config

const (
	ITEM_POKEBALL = "Pokeball"
	ITEM_POKEDEX  = "Pokedex"
)

func getPokeballInfo() ItemBaseInfo {
	return ItemBaseInfo{
		Name:        ITEM_POKEBALL,
		Price:       100,
		Description: "A device for catching wild Pokémon. It is thrown like a ball at a Pokémon, comfortably encapsulating its target.",
	}
}

func getPokedexInfo() ItemBaseInfo {
	return ItemBaseInfo{
		Name:        ITEM_POKEDEX,
		Price:       0,
		Description: "A high-tech device that records the Pokémon you've seen and caught. It's a high-tech encyclopedia!",
	}
}

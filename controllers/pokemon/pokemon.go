package pokemonctr

import "pokemon-game-api/domains/models"

type PokemonControllerResponse struct {
	Id        int                  `json:"id"`
	Name      string               `json:"name"`
	Height    int                  `json:"height"`
	Weight    int                  `json:"weight"`
	Types     []string             `json:"types"`
	Abilities []string             `json:"abilities"`
	Stats     []models.PokemonStat `json:"stats"`
	Sound     string               `json:"sound"`
	Avatar    models.PokemonAvatar `json:"avatar"`
}

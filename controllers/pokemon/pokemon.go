package pokemonctr

import "pokemon-game-api/domains/models"

// omitempty in json will hide json property in json response if this field have not data
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

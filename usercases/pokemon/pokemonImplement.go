package pokemonusc

import (
	"pokemon-game-api/domains/models"
	pokemongwy "pokemon-game-api/gateways/pokemon"
)

type PokemonUsecase struct {
	PokemonGateway pokemongwy.IPokemonGateway
}

func NewPokemonUsecase(pokemonGateway pokemongwy.IPokemonGateway) IPokemonUsecase {
	return PokemonUsecase{PokemonGateway: pokemonGateway}
}

func (p PokemonUsecase) GetPokemonDetailById(id string) (*PokemonUsecaseResponse, error) {
	gwyResponse, cErr := p.PokemonGateway.GetPokeapiPokemonDetailById(id)

	if cErr != nil {
		return nil, cErr
	}

	result := PokemonUsecaseResponse{
		Id:     gwyResponse.ID,
		Name:   gwyResponse.Name,
		Height: gwyResponse.Height,
		Weight: gwyResponse.Weight,
		Sound:  gwyResponse.Cries.Latest,
		Avatar: models.PokemonAvatar{
			Static:  gwyResponse.Sprites.Other.OfficialArtwork.FrontDefault,
			Animate: gwyResponse.Sprites.Other.Showdown.FrontDefault,
		},
	}

	for _, v := range gwyResponse.Types {
		result.Types = append(result.Types, v.Type.Name)
	}

	for _, v := range gwyResponse.Abilities {
		result.Abilities = append(result.Types, v.Ability.Name)
	}

	for _, v := range gwyResponse.Stats {
		result.Stats = append(result.Stats, models.PokemonStat{
			Name:   v.Stat.Name,
			Value:  v.BaseStat,
			Effort: v.Effort,
		})
	}

	return &result, nil
}

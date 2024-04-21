package utils

import (
	"pokemon-game-api/domains/models"
	pokemonconst "pokemon-game-api/pkgs/constants/pokemon"
	"strconv"
)

var regions = [6]models.Region{
	{
		Name:   pokemonconst.NationalName,
		Number: pokemonconst.NationalNumber,
	},
	{
		Name:   pokemonconst.KantoName,
		Number: pokemonconst.KantoNumber,
	},
	{
		Name:   pokemonconst.JohtoName,
		Number: pokemonconst.JohtoNumber,
	},
	{
		Name:   pokemonconst.HoennName,
		Number: pokemonconst.HoennNumber,
	},
	{
		Name:   pokemonconst.SinnohName,
		Number: pokemonconst.SinnohNumber,
	},
	{
		Name:   pokemonconst.OgSinnohName,
		Number: pokemonconst.OgSinnohNumber,
	},
}

func GetRegionNo(region string) string {
	for _, v := range regions {
		if v.Name == region {
			return strconv.Itoa(v.Number)
		}
	}

	return strconv.Itoa(GetDefaultRegion().Number)
}

func GetDefaultRegion() models.Region {
	return regions[0]
}

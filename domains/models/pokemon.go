package models

type PokemonStat struct {
	Name   string `json:"name"`
	Value  int    `json:"value"`
	Effort int    `json:"effort"`
}

type PokemonAvatar struct {
	Static  string `json:"static"`
	Animate string `json:"animate"`
}

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

type Region struct {
	Name   string
	Number int
}

type Move struct {
	Id       int    `json:"id"`
	Sequence int    `json:"squence"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

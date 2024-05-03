package authctr

import (
	"pokemon-game-api/domains/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterRequest struct {
	UserName    string `json:"user_name"`
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
}

// omitempty in json will hide json property in json response if this field have not data
type UserResponse struct {
	Id          primitive.ObjectID `json:"id"`
	UserName    string             `json:"user_name"`
	DisplayName string             `json:"display_name"`
	Password    string             `json:"password"`
	IvKey       string             `json:"iv_key"`
	CreateAt    uint32             `json:"create_at"`
	Pokemons    []PokemonDto       `json:"pokemons"`
}

type PokemonDto struct {
	Id       primitive.ObjectID `json:"id" `
	Name     string             `json:"name"`
	Nickname string             `json:"nickname"`
	Level    string             `json:"lv"`
	Sequence int                `json:"squence"`
	Moves    []models.Move      `json:"moves"`
}

type LoginRequest struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

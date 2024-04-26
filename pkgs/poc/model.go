package poc

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name     Name                 `json:"name,omitempty" bson:"name,inline,omitempty"`
	Info     Info                 `json:"info,omitempty" bson:"info,omitempty"`
	Pokemons []primitive.ObjectID `json:"pokemons,omitempty" bson:"pokemon,omitempty"`
}

type Pokemon struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Moves []string           `json:"moves,omitempty" bson:"moves,omitempty"`
	User  User               `json:"user,omitempty" bson:"user,omitempty"`
}

type Info struct {
	Age   int    `json:"age,omitempty" bson:"age,omitempty"`
	Color string `json:"color,omitempty" bson:"color,omitempty"`
}

type Name struct {
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
}

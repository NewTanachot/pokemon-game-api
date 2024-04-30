package poc

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name       Name                 `json:"name" bson:"name,inline,omitempty"`
	Info       Info                 `json:"info" bson:"info,omitempty"`
	PokemonIds []primitive.ObjectID `json:"pokemon_ids" bson:"pokemon_ids,omitempty"`
	Pokemons   []Pokemon            `json:"pokemons_lookup" bson:"pokemons_lookup,omitempty"`
}

type Pokemon struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name,omitempty"`
	Moves []string           `json:"moves" bson:"moves,omitempty"`
	// UserId primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`

	/* -=-=-=-=-=-=-=-=- [ NOTE ] -=-=-=-=-=-=-=-=-

	- when use Struct or non default type (not pointer) if body of request did not have user object. it will be empty object but "it not nil"
	so when it store in mongo it still have object in there even "omitempty"
	- to prevent that, use *Struct (pointer) and it will be absolute nil
	or another way is check struct when it parse to struct before save it to mongo

	*/
}

type Info struct {
	Age   int    `json:"age" bson:"age"`
	Color string `json:"color" bson:"color,omitempty"`
}

type Name struct {
	FirstName string `json:"first_name" bson:"first_name,omitempty"`
	LastName  string `json:"last_name" bson:"last_name,omitempty"`
}

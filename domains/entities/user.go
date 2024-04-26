package entities

// omitempty in entity will hide bson property in bson in mongo database if this field have no data
type User struct {
	Id          string    `bson:"_id,omitempty"`
	UserName    string    `bson:"user_name,omitempty"`
	DisplayName string    `bson:"display_name,omitempty"`
	Password    string    `bson:"password,omitempty"`
	IvKey       string    `bson:"iv_key,omitempty"`
	Pokemons    []Pokemon `bson:"pokemons,omitempty"`
}

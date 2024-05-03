package entities

import (
	"pokemon-game-api/domains/models"
)

// omitempty in entity will hide bson property in bson in mongo database if this field have no data
type Pokemon struct {
	BaseEntity `bson:"inline,omitempty"`
	Name       string        `bson:"name,omitempty"`
	Nickname   string        `bson:"nickname,omitempty"`
	Level      string        `bson:"lv,omitempty"`
	Sequence   int           `bson:"sequence,omitempty"`
	Moves      []models.Move `bson:"moves,omitempty"`
}

package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseEntity struct {
	// id must be omit empty, because if it not omit it will send default value '0000000' to database when insert entity
	Id       primitive.ObjectID  `bson:"_id,omitempty"`
	CreateAt primitive.Timestamp `bson:"create_at,omitempty"`
}

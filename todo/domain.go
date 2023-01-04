package todo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Domain of Todo
type Domain struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Task      string             `bson:"task,omitempty" json:"task,omitempty"`
	Completed bool               `bson:"completed" json:"completed"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

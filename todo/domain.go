package todo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Domain of Todo
type Domain struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Task      string             `bson:"task" json:"task"`
	Done      bool               `bson:"done" json:"done"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResourceLink struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Link string             `json:"link" bson:"link"`
}

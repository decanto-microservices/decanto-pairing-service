package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Winefamily struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	Winefamily string             `bson:"winefamily" json:"winefamily"`
}

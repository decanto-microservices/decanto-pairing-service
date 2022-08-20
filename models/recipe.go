package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id"`
	Name            string             `bson:"name" json:"name"`
	IsRedAffine     bool               `bson:"_isRedAffine" json:"_isRedAffine"`
	IsWhiteAffine   bool               `bson:"_isWhiteAffine" json:"_isWhiteAffine"`
	IsSparkAffine   bool               `bson:"_isSparkAffine" json:"_isSparkAffine"`
	IsSweetAffine   bool               `bson:"_isSweetAffine" json:"_isSweetAffine"`
	IsLiqueurAffine bool               `bson:"_isLiqueurAffine" json:"_isLiqueurAffine"`
	IsRoseAffine    bool               `bson:"_isRoseAffine" json:"_isRoseAffine"`
	Structure       float64            `bson:"_rStructure" json:"_rStructure"`
	Softness        float64            `bson:"_rSoftness" json:"_rSoftness"`
	Hardness        float64            `bson:"_rHardness" json:"_rHardness"`
	Sweetness       float64            `bson:"v_Dolcezza" json:"v_Dolcezza"`
	FoodSx          float64            `bson:"_foodSx" json:"_foodSx"`
	FoodDx          float64            `bson:"_foodDx" json:"_foodDx"`
}

package models

/*
This struct holds organoleptic characteristics about food, they will be used to build the data
which will be used to query winefamilies
*/
type Food struct {
	Structure float64 `bson:"_rStructure" json:"_rStructure"`
	Softness  float64 `bson:"_rSoftness" json:"_rSoftness"`
	Hardness  float64 `bson:"_rHardness" json:"_rHardness"`
	Sweetness float64 `bson:"v_Dolcezza" json:"v_Dolcezza"`
	FoodSx    float64 `bson:"_foodSx" json:"_foodSx"`
	FoodDx    float64 `bson:"_foodDx" json:"_foodDx"`
}

func NewFood(recipe Recipe) *Food {
	food := &Food{}

	food.Structure = recipe.Structure
	food.Softness = recipe.Softness
	food.Hardness = recipe.Hardness
	food.Sweetness = recipe.Sweetness
	food.FoodSx = recipe.FoodSx
	food.FoodDx = recipe.FoodDx

	return food
}

func (Food) StructureDelta() float64 {
	return 1.0
}

func (Food) SoftnessDelta() float64 {
	return 1.5
}

func (Food) HardnessDelta() float64 {
	return 1.5
}

func (f Food) SweetnessDelta() float64 {
	var sweetness float64 = 1.0

	if f.Sweetness < 4 {
		sweetness = 4
	}

	return sweetness
}

func (Food) DXDelta() float64 {
	return 1.5
}

func (Food) SXDelta() float64 {
	return 1.5
}

package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gprisco/decanto-pairing-service/consul"
	"github.com/Gprisco/decanto-pairing-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PairFamilies(recipeId primitive.ObjectID) models.Recipe { // Temporaneamente ritornando models.Recipe
	servicesMap := consul.Discovery()
	foodService := servicesMap["decanto-food-service"]

	resp, err := http.Get(fmt.Sprintf("http://%s:%v/decanto/food/recipe/%s", foodService.Address, foodService.Port, recipeId.Hex()))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var recipe models.Recipe
	err = json.NewDecoder(resp.Body).Decode(&recipe)

	if err != nil {
		panic(err)
	}

	if recipe.IsLiqueurAffine {
		println("Affine ai liquorosi")
	}

	if recipe.IsRedAffine {
		println("Affine ai rossi")
	}

	if recipe.IsRoseAffine {
		println("Affine ai rose")
	}

	if recipe.IsSparkAffine {
		println("Affine ai frizzanti")
	}

	if recipe.IsSweetAffine {
		println("Affine ai dolci")
	}

	if recipe.IsWhiteAffine {
		println("Affine ai bianchi")
	}

	return recipe
}

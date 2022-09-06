package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gprisco/decanto-pairing-service/env"
	"github.com/Gprisco/decanto-pairing-service/helpers"
	"github.com/Gprisco/decanto-pairing-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PairFamilies(recipeId primitive.ObjectID) []models.Winefamily { // Temporaneamente ritornando models.Recipe
	resp, err := http.Get(fmt.Sprintf("http://%s/decanto/food/recipe/%s", env.GetInstance().FoodURL, recipeId.Hex()))

	if err != nil {
		println(err.Error())
		return []models.Winefamily{}
	}

	defer resp.Body.Close()

	var recipe models.Recipe
	err = json.NewDecoder(resp.Body).Decode(&recipe)

	if err != nil {
		panic(err)
	}

	winetypeIds, winecolorIds := getWineTypesAndColors(recipe)

	food := models.NewFood(recipe)

	winefamilies := getFamilies(winetypeIds, winecolorIds, *food)

	return winefamilies
}

func getWineTypesAndColors(recipe models.Recipe) (winetypeIds []int, winecolorIds []int) {
	winetypeIds = make([]int, 6)
	winecolorIds = make([]int, 4)

	if recipe.IsLiqueurAffine {
		winecolorIds = append(winecolorIds, 1, 2, 3)
		winetypeIds = append(winetypeIds, 6)
	}

	if recipe.IsRedAffine {
		winecolorIds = append(winecolorIds, 1)
		winetypeIds = append(winetypeIds, 2, 5)
	}

	if recipe.IsRoseAffine {
		winecolorIds = append(winecolorIds, 3)
		winetypeIds = append(winetypeIds, 2, 5)
	}

	if recipe.IsSparkAffine {
		winecolorIds = append(winecolorIds, 1, 2, 3)
		winetypeIds = append(winetypeIds, 3, 4)
	}

	if recipe.IsSweetAffine {
		winecolorIds = append(winecolorIds, 1, 2, 3)
		winetypeIds = append(winetypeIds, 1)
	}

	if recipe.IsWhiteAffine {
		winecolorIds = append(winecolorIds, 2)
		winetypeIds = append(winetypeIds, 2, 5)
	}

	winecolorIds = helpers.RemoveDuplicate(winecolorIds)
	winetypeIds = helpers.RemoveDuplicate(winetypeIds)

	return
}

func getFamilies(winetypeIds []int, winecolorIds []int, food models.Food) []models.Winefamily {
	queryString := "?page=1&limit=50&winetypeIds="

	for index, elem := range winetypeIds {
		queryString += fmt.Sprintf("%d", elem)

		if index < len(winetypeIds)-1 {
			queryString += ","
		}
	}

	queryString += "&winecolorIds="

	for index, elem := range winecolorIds {
		queryString += fmt.Sprintf("%d", elem)

		if index < len(winecolorIds)-1 {
			queryString += ","
		}
	}

	queryString += fmt.Sprintf("&structure=%.2f", food.Structure)
	queryString += fmt.Sprintf("&structureD=%.2f", food.StructureDelta())

	queryString += fmt.Sprintf("&softness=%.2f", food.Softness)
	queryString += fmt.Sprintf("&softnessD=%.2f", food.SoftnessDelta())

	queryString += fmt.Sprintf("&hardness=%.2f", food.Hardness)
	queryString += fmt.Sprintf("&hardnessD=%.2f", food.HardnessDelta())

	queryString += fmt.Sprintf("&sweetness=%.2f", food.Sweetness)
	queryString += fmt.Sprintf("&sweetnessD=%.2f", food.SweetnessDelta())

	queryString += fmt.Sprintf("&foodSx=%.2f", food.FoodSx)
	queryString += fmt.Sprintf("&foodSxD=%.2f", food.SXDelta())

	queryString += fmt.Sprintf("&foodDx=%.2f", food.FoodDx)
	queryString += fmt.Sprintf("&foodDxD=%.2f", food.DXDelta())

	var winefamilies []models.Winefamily

	resp, err := http.Get(fmt.Sprintf("http://%s/decanto/winefamily"+queryString, env.GetInstance().WinefamilyURL))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&winefamilies)

	if err != nil {
		return nil
	}

	return winefamilies
}

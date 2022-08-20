package handlers

import (
	"net/http"

	"github.com/Gprisco/decanto-pairing-service/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PairFamilies(c *gin.Context) {
	recipeId, err := primitive.ObjectIDFromHex(c.Param("recipeId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "This is not a valid ID")
		return
	}

	recipe := services.PairFamilies(recipeId)
	c.JSON(http.StatusOK, recipe)
}

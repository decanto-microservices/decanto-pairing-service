package main

import (
	"net/http"

	"github.com/Gprisco/decanto-pairing-service/consul"
	"github.com/Gprisco/decanto-pairing-service/env"
	"github.com/Gprisco/decanto-pairing-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	consul.Register()

	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET(baseUrl+"/check", (func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}))

	r.GET(baseUrl+"/:recipeId", handlers.PairFamilies)

	r.Run(env.GetInstance().Port)
}

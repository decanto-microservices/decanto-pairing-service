package main

import (
	"net/http"

	"github.com/Gprisco/decanto-pairing-service/env"
	"github.com/Gprisco/decanto-pairing-service/services"
	"github.com/gin-gonic/gin"
)

func main() {
	services.Register()
	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET(baseUrl+"/check", (func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}))

	r.Run(env.GetInstance().Port)
}

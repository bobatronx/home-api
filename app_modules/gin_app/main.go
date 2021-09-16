package main

import (
	"net/http"

	"github.com/bobatronx/gin-app/v2/models"
	"github.com/bobatronx/gin-app/v2/services"
	"github.com/gin-gonic/gin"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

func main() {
	router := gin.Default()
	lands := router.Group("/v1")
	{
		lands.GET("/lands", getLandsEndpoint)
		lands.POST("/lands", addLandsEndpoint)
	}

	router.Run(":9000")
}

func getLandsEndpoint(c *gin.Context) {
	lands := services.GetLands()
	c.JSON(http.StatusOK, lands)
}

func addLandsEndpoint(c *gin.Context) {
	land := new(models.Land)
	if err := c.ShouldBindJSON(land); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created_land, err := services.AddLand(land)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created_land)
}

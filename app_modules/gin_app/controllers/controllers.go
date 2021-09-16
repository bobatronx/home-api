package controllers

import (
	"net/http"

	"github.com/bobatronx/gin-app/v2/models"
	"github.com/bobatronx/gin-app/v2/services"
	"github.com/gin-gonic/gin"
)

type landRequest struct {
	Acres     string            `json:"acres"`
	Buildings []buildingRequest `json:"buildings"`
}

type buildingRequest struct {
	Name string `json:"name"`
}

func InitializeControllers(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/lands", registerGetLandEndpoint)
		v1.POST("/lands", registerAddLandEndpoint)
	}
}

func registerGetLandEndpoint(c *gin.Context) {
	lands := services.GetLands()
	c.JSON(http.StatusOK, lands)
}

func registerAddLandEndpoint(c *gin.Context) {
	land := new(landRequest)
	if err := c.ShouldBindJSON(land); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buildings := []*models.Building{}

	for _, v := range land.Buildings {
		building := models.Building(v)
		buildings = append(buildings, &building)
	}

	new_land := models.Land{
		Acres:     land.Acres,
		Buildings: buildings,
	}

	created_land, err := services.AddLand(&new_land)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created_land)
}

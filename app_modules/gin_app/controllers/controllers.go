package controllers

import (
	"net/http"
	"strconv"

	"github.com/bobatronx/gin-app/v2/models"
	"github.com/bobatronx/gin-app/v2/service"
	"github.com/gin-gonic/gin"
)

func InitializeControllers(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/lands", getLandsController)
		v1.POST("/lands", createLandController)
		v1.GET("/lands/:id", getLandByIdController)
		v1.PUT("/lands/:id", updateLandController)
		v1.DELETE("/lands/:id", deleteLandController)
	}
}

func getLandsController(c *gin.Context) {
	lands := service.GetLands()
	c.JSON(http.StatusOK, lands)
}

func getLandByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Land id must be a valid integer"})
		return
	}

	land := service.GetLandById(id)

	if land.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, land)
}

func createLandController(c *gin.Context) {

	land := models.Land{}

	if err := c.ShouldBindJSON(&land); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLand := service.CreateLand(&land)

	c.JSON(http.StatusCreated, createdLand)
}

func updateLandController(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Land id must be a valid integer"})
		return
	}

	land := models.Land{ID: id}

	if err := c.ShouldBindJSON(&land); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	land.ID = id

	updatedLand := service.UpdateLand(&land)

	c.JSON(http.StatusOK, updatedLand)
}

func deleteLandController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Land id must be a valid integer"})
		return
	}

	success := service.DeleteLandById(id)

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete the specified land"})
		return
	}

	c.Status(http.StatusOK)
}

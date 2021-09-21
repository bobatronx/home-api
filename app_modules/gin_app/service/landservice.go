package service

import "github.com/bobatronx/gin-app/v2/models"
import "github.com/bobatronx/gin-app/v2/repository"


func GetLands() []models.Land {
	return repository.GetLands()
}

func GetLandById(id int) models.Land {
	return repository.GetLandById(id)
}

func CreateLand(land *models.Land) models.Land {
	return repository.CreateLand(land)
}

func UpdateLand(land *models.Land) models.Land {
	return repository.UpdateLand(land)
}

func DeleteLandById(id int) bool {
	return repository.DeleteLandById(id)
}

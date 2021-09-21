package repository

import "github.com/bobatronx/gin-app/v2/models"

var (
	lands  []*models.Land
	nextID = 1
)

func CreateLand(land *models.Land) models.Land {
	land.ID = nextID
	lands = append(lands, land)
	nextID++

	return *land
}

func UpdateLand(landUpdate *models.Land) models.Land {
	for _, land := range lands {
		if land.ID == landUpdate.ID {
			land.Acres = landUpdate.Acres
			land.Buildings = landUpdate.Buildings

			return *land
		}
	}

	return models.Land{}
}

func GetLandById(id int) models.Land {
	for _, land := range lands {
		if land.ID == id {
			return *land
		}
	}

	return models.Land{}
}

func GetLands() []models.Land {
	landsDisplay := []models.Land{}

	for _, land := range lands {
		landsDisplay = append(landsDisplay, *land)
	}

	return landsDisplay
}

func DeleteLandById(id int) bool {
	for i, land := range lands {
		if land.ID == id {
			lands = append(lands[:i], lands[i+1:]...)
			return true
		}
	}

	return false
}

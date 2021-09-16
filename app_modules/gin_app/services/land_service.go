package services

import "github.com/bobatronx/gin-app/v2/models"

var (
	lands  []*models.Land
	nextID int
)

func GetLands() []*models.Land {
	return lands
}

func AddLand(land *models.Land) (*models.Land, error) {
	land.ID = nextID
	lands = append(lands, land)

	nextID++

	return land, nil
}

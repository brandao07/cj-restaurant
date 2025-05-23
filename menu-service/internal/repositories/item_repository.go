package repositories

import (
	"github.com/brandao07/cj-restaurant/menu-service/internal/data/models"
)

type ItemRepository interface {
	GetItemByID(id int) (models.Item, error)
	GetItemsByIDs(ids []int) ([]models.Item, error)
}

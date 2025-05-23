package repositories

import (
	"github.com/brandao07/cj-restaurant/menu-service/internal/dtos"
)

type ItemRepository interface {
	GetItemByID(id int) (dtos.Item, error)
	GetItemsByIDs(ids []int) ([]dtos.Item, error)
}

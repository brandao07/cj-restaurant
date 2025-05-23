package repositories

import (
	"github.com/brandao07/cj-restaurant/menu-service/internal/data/models"
	"github.com/brandao07/cj-restaurant/menu-service/internal/dtos"
)

type MenuRepository interface {
	CreateMenu(CreateMenuDTO *dtos.CreateMenuDTO) (*models.Menu, error)
	DeactivateMenu(id int) error
}

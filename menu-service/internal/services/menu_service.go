package services

import (
	"fmt"

	"github.com/brandao07/cj-restaurant/menu-service/internal/common"
	"github.com/brandao07/cj-restaurant/menu-service/internal/dtos"
	"github.com/brandao07/cj-restaurant/menu-service/internal/repositories"
)

type MenuService struct {
	MenuRepository repositories.MenuRepository
	ItemRepository repositories.ItemRepository
}

func NewMenuService(menuRepository repositories.MenuRepository, itemRepository repositories.ItemRepository) *MenuService {
	return &MenuService{
		MenuRepository: menuRepository,
		ItemRepository: itemRepository,
	}
}

func (s *MenuService) CreateMenu(createMenuDTO *dtos.CreateMenuDTO) (*dtos.Menu, error) {
	// Fetch the items by their IDs
	itemIDs := []int{createMenuDTO.StarterID, createMenuDTO.MainID, createMenuDTO.DessertID, createMenuDTO.DrinkID}
	items, err := s.ItemRepository.GetItemsByIDs(itemIDs)
	if err != nil {
		return nil, err
	} else if len(items) != 4 {
		return nil, fmt.Errorf("not all items found: expected 4, got %d", len(items))
	}

	

	// Map item IDs to names
	idToName := make(map[int]string)
	for _, item := range items {
		idToName[item.ID] = item.Name
	}

	menu, err := s.MenuRepository.CreateMenu(createMenuDTO)
	if err != nil {
		return nil, err
	}

	return &dtos.Menu{
		ID: menu.ID,
		Items: map[string]string{
			common.MenuStarter: idToName[menu.StarterID],
			common.MenuMain:    idToName[menu.MainID],
			common.MenuDessert: idToName[menu.DessertID],
			common.MenuDrink:   idToName[menu.DrinkID],
		},
		Price: menu.Price,
	}, nil
}

func (s *MenuService) DeactivateMenu(id int) error {
	err := s.MenuRepository.DeactivateMenu(id)
	if err != nil {
		return err
	}

	return nil
}

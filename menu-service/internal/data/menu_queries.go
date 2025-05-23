package data

import (
	"fmt"

	"github.com/brandao07/cj-restaurant/menu-service/internal/common"
	"github.com/brandao07/cj-restaurant/menu-service/internal/data/models"
	"github.com/brandao07/cj-restaurant/menu-service/internal/dtos"
)

func (db *DB) CreateMenu(dto *dtos.CreateMenuDTO) (*models.Menu, error) {
	query := `
        INSERT INTO menus (starter_id, main_id, dessert_id, drink_id, price)
        VALUES (:starter_id, :main_id, :dessert_id, :drink_id, :price)
        RETURNING id
    `

	menu := models.Menu{
		StarterID: dto.StarterID,
		MainID:    dto.MainID,
		DessertID: dto.DessertID,
		DrinkID:   dto.DrinkID,
		Price:     dto.Price,
	}
	nstmt, err := db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer nstmt.Close()

	if err := nstmt.Get(&menu, menu); err != nil {
		return nil, err
	}
	return &menu, nil
}

func (db *DB) DeactivateMenu(id int) error {
	query := `
		UPDATE menus
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("0 rows affected, menu with id %d may not exist or is already deleted", id)
	}

	return nil
}

func (db *DB) GetAllActiveMenus() ([]dtos.Menu, error) {
	query := `
		SELECT
			m.id,
			s.name AS starter_name,
			ma.name AS main_name,
			d.name AS dessert_name,
			dr.name AS drink_name,
			m.price
		FROM menus m
		LEFT JOIN items s ON m.starter_id = s.id
		LEFT JOIN items ma ON m.main_id = ma.id
		LEFT JOIN items d ON m.dessert_id = d.id
		LEFT JOIN items dr ON m.drink_id = dr.id
		WHERE m.deleted_at IS NULL
	`

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menus []dtos.Menu
	for rows.Next() {
		var menu models.MenuWithNames
		if err := rows.StructScan(&menu); err != nil {
			return nil, err
		}
		items := map[string]string{
			common.MenuStarter: menu.StarterName,
			common.MenuMain:    menu.MainName,
			common.MenuDessert: menu.DessertName,
			common.MenuDrink:   menu.DrinkName,
		}
		menus = append(menus, *dtos.NewMenu(menu.ID, items, menu.Price))
	}
	return menus, nil
}

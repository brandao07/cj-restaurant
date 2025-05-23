package data

import (
	"fmt"
	"strings"

	"github.com/brandao07/cj-restaurant/menu-service/internal/data/models"
	"github.com/brandao07/cj-restaurant/menu-service/internal/dtos"
)

func (db *DB) GetItemByID(id int) (dtos.Item, error) {
	panic("unimplemented")
}

func (db *DB) GetItemsByIDs(ids []int) ([]dtos.Item, error) {
	if len(ids) == 0 {
		return []dtos.Item{}, nil
	}

	// Build the query with the correct number of placeholders
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))

	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1) // $1, $2, ...
		args[i] = id
	}

	// Dynamic query because the number of IDs can change
	query := fmt.Sprintf(
		`SELECT id, name, type FROM items WHERE id IN (%s)`,
		strings.Join(placeholders, ","), // Join the list of placeholders with ','
	)

	rows, err := db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []dtos.Item
	for rows.Next() {
		var item models.Item
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}
		items = append(items, *dtos.NewItem(item.ID, item.Name, item.Type))
	}
	return items, nil
}

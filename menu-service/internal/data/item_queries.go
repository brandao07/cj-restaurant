package data

import (
	"fmt"
	"strings"

	"github.com/brandao07/cj-restaurant/menu-service/internal/data/models"
)

func (db *DB) GetItemByID(id int) (models.Item, error) {
	panic("unimplemented")
}

func (db *DB) GetItemsByIDs(ids []int) ([]models.Item, error) {
    if len(ids) == 0 {
        return []models.Item{}, nil
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

    var items []models.Item
    for rows.Next() {
        var item models.Item
        if err := rows.StructScan(&item); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}
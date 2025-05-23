package dtos

// Item represents a menu response
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewItem(id int, name string, itemType string) *Item {
	return &Item{
		ID:   id,
		Name: name,
		Type: itemType,
	}
}

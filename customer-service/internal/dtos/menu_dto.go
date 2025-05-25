package dtos

// Menu represents a menu response
type Menu struct {
	ID    int               `json:"id"`
	Items map[string]string `json:"items"`
	Price float32           `json:"price"`
}

func NewMenu(id int, items map[string]string, price float32) *Menu {
	return &Menu{
		ID:    id,
		Items: items,
		Price: price,
	}
}

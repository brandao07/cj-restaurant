package dtos

type CreateOrderDTO struct {
	MenuIDs    []int   `json:"menu_ids"`
	Total      float64 `json:"total"`
	Status     string  `json:"status"`
}

func NewCreateOrderDTO(menuIDs []int, total float64, status string) *CreateOrderDTO {
	return &CreateOrderDTO{
		MenuIDs:    menuIDs,
		Total:      total,
		Status:     status,
	}
}

// TODO: MoreDTOs..

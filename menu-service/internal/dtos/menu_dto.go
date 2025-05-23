package dtos

// Input DTOs 

type CreateMenuDTO struct {
    StarterID int  `json:"starter"`
    MainID    int  `json:"main"`
    DessertID int  `json:"dessert"`
    DrinkID   int  `json:"drink"`
    Price     float32 `json:"price"`
}

func NewCreateMenuDTO(starterID, mainID, dessertID, drinkID int, price float32) *CreateMenuDTO {
    return &CreateMenuDTO{
        StarterID: starterID,
        MainID:    mainID,
        DessertID: dessertID,
        DrinkID:   drinkID,
        Price:   price,
    }
}

// Output DTOs

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
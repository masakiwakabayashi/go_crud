package dto

type CreateItemInput struct {
	Name        string `json:"name" biding:"required,min=2"`
	Price       uint   `json:"price" biding:"required,min=1,max=99999"`
	Description string `json:"description"`
}

type UpdateItemInput struct {
	Name        *string `json:"name" biding:"omitnil,min=2"`
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}

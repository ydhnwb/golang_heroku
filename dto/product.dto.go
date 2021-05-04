package dto

type CreateProductRequest struct {
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Price uint64 `json:"price" form:"email" binding:"required"`
}

type UpdateProductRequest struct {
	ID    int64  `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Price uint64 `json:"price" form:"email" binding:"required"`
}

package dto

type ID struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

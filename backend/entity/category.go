package entity

type Category struct {
    Title string `json:"title" binding:"required"`
}

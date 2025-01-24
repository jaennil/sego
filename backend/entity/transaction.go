package entity

import "time"

type Transaction struct {
    ID        uint      `json:"id"`
    Amount    float64   `json:"amount" binding:"required"`
    Type      string    `json:"type" binding:"required"`
    CreatedAt time.Time `json:"created_at" binding:"required"`
    Account   string    `json:"account" binding:"required"`
    Category  string    `json:"category" binding:"required"`
}

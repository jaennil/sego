package entity

import "time"

type Account struct {
    Name      string    `json:"name" binding:"required"`
    Type      string    `json:"type" binding:"required"`
    Balance   float64   `json:"balance" binding:"required"`
    CreatedAt time.Time `json:"created_at" binding:"required"`
    Currency  string    `json:"currency" binding:"required"`
}

package models

import "time"


type Order struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	UserName    string    `json:"user_name"`
	ProductName string    `json:"product_name"`
	Quantity    uint      `json:"quantity"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderRequest struct {
	UserID      uint        `json:"userId"`
	ProductName string      `json:"product_name"`
	Quantity    uint        `json:"quantity"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

const (
	Pending    = "Pending"
	Processing = "Processing"
	Shipped    = "Shipped"
	Delivered  = "Delivered"
	Completed  = "Completed"
	Canceled   = "Canceled"
	Refunded   = "Refunded"
	Failed     = "Failed"
)

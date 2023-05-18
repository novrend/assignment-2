package models

import "time"

type OrderResponse struct {
	OrderId      uint           `gorm:"column:order_id;primary_key" json:"id"`
	CustomerName string         `gorm:"column:customer_name;type:varchar(100)" json:"customer_name"`
	Items        []ItemResponse `gorm:"foreignKey:order_id" json:"Items"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

func (OrderResponse) TableName() string {
	return "orders"
}

type Order struct {
	OrderId      uint      `gorm:"column:order_id;primary_key" json:"-"`
	CustomerName string    `gorm:"column:customer_name;type:varchar(100)" json:"customerName"`
	Items        []Item    `gorm:"foreignKey:order_id" json:"items"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"orderedAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"-"`
}

type OrderUpdateResponse struct {
	OrderId      uint                 `gorm:"column:order_id;primary_key" json:"orderId"`
	CustomerName string               `gorm:"column:customer_name;type:varchar(100)" json:"customerName"`
	Items        []ItemUpdateResponse `gorm:"foreignKey:order_id" json:"items"`
	CreatedAt    time.Time            `gorm:"column:created_at" json:"orderedAt"`
	UpdatedAt    time.Time            `gorm:"column:updated_at" json:"-"`
}

func (OrderUpdateResponse) TableName() string {
	return "orders"
}

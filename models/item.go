package models

import "time"

type ItemResponse struct {
	ItemId      uint      `gorm:"column:item_id;primary_key" json:"id"`
	ItemCode    string    `gorm:"column:item_code;type:varchar(100)" json:"item_code"`
	Description string    `gorm:"varchar(100)" json:"description"`
	Quantity    uint      `gorm:"integer(5)" json:"quantity"`
	OrderId     uint      `gorm:"column:order_id" json:"order_id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ItemResponse) TableName() string {
	return "items"
}

type Item struct {
	ItemId      uint      `gorm:"column:item_id;primary_key" json:"-"`
	ItemCode    string    `gorm:"column:item_code;type:varchar(100)" json:"itemCode"`
	Description string    `gorm:"varchar(100)" json:"description"`
	Quantity    uint      `gorm:"integer(5)" json:"quantity"`
	OrderId     uint      `gorm:"column:order_id" json:"-"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-"`
}

type ItemUpdateResponse struct {
	ItemId      uint      `gorm:"column:item_id;primary_key" json:"lineItemId"`
	ItemCode    string    `gorm:"column:item_code;type:varchar(100)" json:"itemCode"`
	Description string    `gorm:"varchar(100)" json:"description"`
	Quantity    uint      `gorm:"integer(5)" json:"quantity"`
	OrderId     uint      `gorm:"column:order_id" json:"-"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-"`
}

func (ItemUpdateResponse) TableName() string {
	return "items"
}

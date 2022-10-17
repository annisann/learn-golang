package structs

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uint32 `json:"item_id" gorm:"primary_key;autoIncrement"`
	ItemCode    string `json:"item_code" gorm:"not null; type:varchar(20)"`
	Description string `json:"description" gorm:"not null; type:varchar(500)"`
	Quantity    uint32 `json:"quantity" gorm:"not null"`
	OrderID     uint32 `json:"order_id" gorm:"not null"`
}

type Order struct {
	gorm.Model
	ID           uint32    `json:"order_id" gorm:"primaryKey; autoIncrement"`
	CustomerName string    `json:"customer_name" gorm:"not null; type:varchar(191)"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID; references:ID"`
}

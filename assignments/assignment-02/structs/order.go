package structs

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint32    `json:"order_id" gorm:"primaryKey; autoIncrement"`
	CustomerName string    `json:"customer_name" gorm:"not null; type:varchar(191)"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID; references:ID"`
}

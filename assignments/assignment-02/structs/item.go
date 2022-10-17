package structs

import (
	"errors"

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

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	if len(item.Description) < 5 {
		err = errors.New("description is too short")
	}
	return
}

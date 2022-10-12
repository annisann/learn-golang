package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

//  association 1:N with user

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null; type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

/* Hooks */
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) { // will be executed before creating data
	fmt.Println("Product before Create()")

	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}
	return
}

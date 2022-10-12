package models

import "time"

//  association 1:N with user

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null; type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

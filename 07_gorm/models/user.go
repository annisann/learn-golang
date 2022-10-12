package models

import "time"

// create table `user` using a struct.

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null; unique;type:varchar(191)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

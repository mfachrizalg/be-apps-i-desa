package models

import "github.com/google/uuid"

type User struct {
	Username  string `gorm:"primaryKey"`
	Password  string
	VillageID uuid.UUID `gorm:"type:uuid;not null"`

	// Belongs to Village
	Village Village `gorm:"foreignKey:VillageID"`
}

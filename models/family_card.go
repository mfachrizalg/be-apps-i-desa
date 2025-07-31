package models

import (
	"github.com/google/uuid"
	"time"
)

type FamilyCard struct {
	NIK           string    `gorm:"primaryKey"`
	Alamat        string    `gorm:"size:50;not null"`
	RT            string    `gorm:"size:50;not null"`
	RW            string    `gorm:"size:50;not null"`
	Kelurahan     string    `gorm:"size:50;not null"`
	Kecamatan     string    `gorm:"size:50;not null"`
	KabupatenKota string    `gorm:"size:50;not null"`
	KodePos       string    `gorm:"size:50;not null"`
	Provinsi      string    `gorm:"size:50;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	VillageID     uuid.UUID `gorm:"type:uuid;not null"`

	// Belongs to Village
	Village Village `gorm:"foreignKey:VillageID"`

	// One-to-many with Villagers
	Villagers []Villager `gorm:"foreignKey:FamilyCardID;references:NIK"`
}

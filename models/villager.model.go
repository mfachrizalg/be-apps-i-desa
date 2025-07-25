package models

import (
	"github.com/google/uuid"
	"time"
)

type Villager struct {
	NIK              string `gorm:"primaryKey"`
	NamaLengkap      string
	JenisKelamin     string
	TempatLahir      string
	TanggalLahir     time.Time
	Agama            string
	Pendidikan       string
	Pekerjaan        string
	StatusPerkawinan string
	StatusHubungan   string
	Kewarganegaraan  string
	NomorPaspor      *string
	NomorKitas       *string
	NamaAyah         string
	NamaIbu          string
	VillageID        uuid.UUID `gorm:"type:uuid;not null"`
	FamilyCardID     string    `gorm:"not null"`

	// Belongs to relationships
	Village    Village    `gorm:"foreignKey:VillageID"`
	FamilyCard FamilyCard `gorm:"foreignKey:FamilyCardID;references:NIK"`
}

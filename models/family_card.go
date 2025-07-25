package models

import "github.com/google/uuid"

type FamilyCard struct {
	NIK           string `gorm:"primaryKey"`
	NamaKepala    string
	Alamat        string
	RT            string
	RW            string
	Kelurahan     string
	Kecamatan     string
	KabupatenKota string
	KodePos       string
	Provinsi      string
	VillageID     uuid.UUID `gorm:"type:uuid;not null"`

	// Belongs to Village
	Village Village `gorm:"foreignKey:VillageID"`

	// One-to-many with Villagers
	Villagers []Villager `gorm:"foreignKey:FamilyCardID;references:NIK"`
}

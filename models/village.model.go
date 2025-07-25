package models

import "github.com/google/uuid"

type Village struct {
	ID                       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Nama                     string
	AksesPaud                string
	AksesSd                  string
	AksesSmp                 string
	AksesSma                 string
	SaranaKesehatan          string
	FasilitasKesehatan       string
	AktivitasPosyandu        string
	LayananDokter            string
	LayananBidan             string
	LayananTenagaKesehatan   string
	JaminanKesehatanNasional string
	AirMinum                 string
	RumahTidakLayakHuni      string

	// One-to-many relationships
	Users       []User       `gorm:"foreignKey:VillageID"`
	Villagers   []Villager   `gorm:"foreignKey:VillageID"`
	FamilyCards []FamilyCard `gorm:"foreignKey:VillageID"`
}

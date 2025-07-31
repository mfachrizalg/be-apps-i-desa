package repositories

import (
	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/models"
	"gorm.io/gorm"
)

type VillageRepository struct {
	DB *gorm.DB
}

func NewVillageRepository() *VillageRepository {
	return &VillageRepository{
		DB: config.DB,
	}
}

func (r *VillageRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *VillageRepository) CreateVillageWithTx(tx *gorm.DB, village *models.Village) error {
	return tx.Create(village).Error
}

package repositories

import (
	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/models"
	"gorm.io/gorm"
)

type VillagerRepository struct {
	DB *gorm.DB
}

func NewVillagerRepository() *VillagerRepository {
	return &VillagerRepository{
		DB: config.DB,
	}
}

func (r *VillagerRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *VillagerRepository) CreateVillagerWithTx(tx *gorm.DB, villager *models.Villager) error {
	return tx.Create(villager).Error
}

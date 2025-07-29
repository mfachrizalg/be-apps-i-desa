package repositories

import (
	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/models"
	"gorm.io/gorm"
)

type FamilyCardRepository struct {
	DB *gorm.DB
}

func NewFamilyCardRepository(db *gorm.DB) *FamilyCardRepository {
	return &FamilyCardRepository{
		DB: config.DB,
	}
}

func (r *FamilyCardRepository) CreateWithTx(tx *gorm.DB, familyCard *models.FamilyCard) error {
	return tx.Create(familyCard).Error
}

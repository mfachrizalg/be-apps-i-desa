package repositories

import (
	"Apps-I_Desa_Backend/config"
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"gorm.io/gorm"
)

type FamilyCardRepository struct {
	DB *gorm.DB
}

func NewFamilyCardRepository() *FamilyCardRepository {
	return &FamilyCardRepository{
		DB: config.DB,
	}
}

func (r *FamilyCardRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *FamilyCardRepository) CreateWithTx(tx *gorm.DB, familyCard *models.FamilyCard) error {
	return tx.Create(familyCard).Error
}

func (r *FamilyCardRepository) GetNIKAndAddressByID(nik string) (*dtos.GetAllFamilyMember, error) {
	var familyCard models.FamilyCard
	err := r.DB.Where("nik = ?", nik).Select("nik", "alamat").First(&familyCard).Error
	if err != nil {
		return nil, err
	}

	return &dtos.GetAllFamilyMember{
		NIK:     familyCard.NIK,
		Address: familyCard.Alamat,
	}, nil
}

func (r *FamilyCardRepository) GetAllFamilyCards() ([]*models.FamilyCard, error) {
	var familyCards []*models.FamilyCard
	err := r.DB.Find(&familyCards).Error
	if err != nil {
		return nil, err
	}
	return familyCards, nil
}

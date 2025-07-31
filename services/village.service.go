package services

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"Apps-I_Desa_Backend/repositories"
	"errors"
	"log"
)

type VillageService struct {
	villageRepo *repositories.VillageRepository
}

func NewVillageService(villageRepo *repositories.VillageRepository) *VillageService {
	return &VillageService{
		villageRepo: villageRepo,
	}
}

func (s *VillageService) CreateVillage(request *dtos.AddVillageRequest) (*dtos.MessageResponse, error) {
	tx := s.villageRepo.BeginTransaction()
	defer tx.Rollback()

	village := &models.Village{
		Name: request.Name,
	}

	if err := s.villageRepo.CreateVillageWithTx(tx, village); err != nil {
		log.Println("Error creating village:", err)
		return nil, errors.New("failed to create village")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		return nil, errors.New("failed to commit transaction")
	}

	return &dtos.MessageResponse{
		Message: "Village created successfully",
	}, nil
}

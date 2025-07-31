package services

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"Apps-I_Desa_Backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

type FamilyCardService struct {
	familyCardRepo *repositories.FamilyCardRepository
}

func NewFamilyCardService(familyCardRepo *repositories.FamilyCardRepository) *FamilyCardService {
	return &FamilyCardService{
		familyCardRepo: familyCardRepo,
	}
}

func (s FamilyCardService) CreateFamilyCard(request *dtos.AddFamilyCardRequest, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	tx := s.familyCardRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village_id").(string)
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Println("Error parsing village ID:", err)
		return nil, errors.New("village ID is not valid")
	}

	familyCard := &models.FamilyCard{
		NIK:           request.NIK,
		Alamat:        request.Address,
		RT:            request.RT,
		RW:            request.RW,
		Kelurahan:     request.Kelurahan,
		Kecamatan:     request.Kecamatan,
		KabupatenKota: request.KabupatenKota,
		Provinsi:      request.Provinsi,
		KodePos:       request.KodePos,
		VillageID:     villageID,
	}

	if err := s.familyCardRepo.CreateWithTx(tx, familyCard); err != nil {
		log.Println("Error creating family card:", err)
		return nil, errors.New("failed to create family card")
	}
	if err := tx.Commit().Error; err != nil {
		log.Println("Error committing transaction:", err)
		return nil, errors.New("failed to create family card")
	}

	return &dtos.MessageResponse{
		Message: "Family card created successfully",
	}, nil
}

func (s FamilyCardService) GetFamilyCardByID(nik string) (*dtos.GetFamilyCardByID, error) {
	response, err := s.familyCardRepo.GetFamilyCardByID(nik)
	if err != nil {
		log.Println("Error getting family card by NIK:", err)
		return nil, errors.New("failed to get family card by NIK")
	}

	if response == nil {
		return nil, errors.New("family card not found")
	}

	return response, nil
}

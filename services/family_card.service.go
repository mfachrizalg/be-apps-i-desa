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
	villagerRepo   *repositories.VillagerRepository
}

func NewFamilyCardService(familyCardRepo *repositories.FamilyCardRepository, villagerRepo *repositories.VillagerRepository) *FamilyCardService {
	return &FamilyCardService{
		familyCardRepo: familyCardRepo,
		villagerRepo:   villagerRepo,
	}
}

func (s *FamilyCardService) CreateFamilyCard(request *dtos.AddFamilyCardRequest, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	tx := s.familyCardRepo.BeginTransaction()
	defer tx.Rollback()

	villageIDStr := ctx.Locals("village")
	if villageIDStr == nil {
		log.Println("Village ID not found in context")
		return nil, errors.New("village ID is required")
	}
	villageID, err := uuid.Parse(villageIDStr.(string))
	// Check if the village ID is valid
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

func (s *FamilyCardService) GetFamilyCardByID(nik string) (*dtos.GetAllFamilyMember, error) {
	response, err := s.familyCardRepo.GetNIKAndAddressByID(nik)
	if err != nil {
		log.Println("Error getting family card by NIK:", err)
		return nil, errors.New("failed to get family card by NIK")
	}

	villagers, err := s.villagerRepo.GetVillagersByFamilyCardID(&nik)
	if err != nil {
		log.Println("Error getting villagers by family card ID:", err)
		return nil, errors.New("failed to get villagers by family card ID")
	}

	if response == nil {
		return nil, errors.New("family card not found")
	}

	// Convert []*dtos.GetFamilyMember to []dtos.GetFamilyMember
	var familyMembers []dtos.GetFamilyMember
	for _, villager := range villagers {
		if villager != nil {
			familyMembers = append(familyMembers, *villager)
		}
	}

	return &dtos.GetAllFamilyMember{
		NIK:           nik,
		Address:       response.Address,
		FamilyMembers: familyMembers,
	}, nil
}

func (s *FamilyCardService) GetAllFamilyCards() (*dtos.GetAllFamilyCardsResponse, error) {
	familyCards, err := s.familyCardRepo.GetAllFamilyCards()
	if err != nil {
		log.Println("Error getting all family cards:", err)
		return nil, errors.New("failed to get all family cards")
	}

	var response dtos.GetAllFamilyCardsResponse
	for _, card := range familyCards {
		villagers, err := s.villagerRepo.GetVillagersByFamilyCardID(&card.NIK)
		if err != nil {
			log.Println("Error getting villagers for family card:", err)
			return nil, errors.New("failed to get villagers for family card")
		}
		var kepalaKeluarga string
		for _, villager := range villagers {
			if villager.StatusHubungan == "Kepala Keluarga" {
				kepalaKeluarga = villager.Name
				break
			}
		}
		response.FamilyCards = append(response.FamilyCards, dtos.GetFamilyCardResponse{
			NIK:          card.NIK,
			Name:         &kepalaKeluarga,
			TotalMembers: len(villagers),
		})
	}

	return &response, nil
}

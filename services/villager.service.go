package services

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"Apps-I_Desa_Backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

type VillagerService struct {
	villagerRepo *repositories.VillagerRepository
}

func NewVillagerService(villagerRepo *repositories.VillagerRepository) *VillagerService {
	return &VillagerService{
		villagerRepo: villagerRepo,
	}
}

func (s *VillagerService) CreateVillager(request *dtos.AddVillagerRequest, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	tx := s.villagerRepo.BeginTransaction()
	defer tx.Rollback()

	// Parse VillageID from string to UUID
	villageIDStr := ctx.Locals("villageID").(string)
	villageID, err := uuid.Parse(villageIDStr)
	if err != nil {
		log.Println("Error parsing village ID:", err)
		return nil, errors.New("invalid village ID format")
	}

	// Parse TanggalLahir from string to time.Time
	tanggalLahir, err := time.Parse("2006-01-02", request.TanggalLahir)
	if err != nil {
		log.Println("Error parsing date:", err)
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	villager := &models.Villager{
		NIK:              request.NIK,
		NamaLengkap:      request.NamaLengkap,
		JenisKelamin:     request.JenisKelamin,
		TempatLahir:      request.TempatLahir,
		TanggalLahir:     tanggalLahir,
		Agama:            request.Agama,
		Pendidikan:       request.Pendidikan,
		Pekerjaan:        request.Pekerjaan,
		StatusPerkawinan: request.StatusPerkawinan,
		StatusHubungan:   request.StatusHubungan,
		Kewarganegaraan:  request.Kewarganegaraan,
		NamaAyah:         request.NamaAyah,
		NamaIbu:          request.NamaIbu,
		VillageID:        villageID,
		FamilyCardID:     request.FamilyCardID,
	}

	if request.NomorKitas != nil {
		villager.NomorKitas = request.NomorKitas
	}

	if request.NomorPaspor != nil {
		villager.NomorPaspor = request.NomorPaspor
	}

	if err := s.villagerRepo.CreateVillagerWithTx(tx, villager); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &dtos.MessageResponse{
		Message: "Villager created successfully",
	}, nil

}

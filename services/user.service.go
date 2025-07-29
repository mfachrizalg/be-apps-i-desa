package services

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"Apps-I_Desa_Backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(request *dtos.RegisterRequest) (*dtos.MessageResponse, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByUsername(request.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Database error: ", err)
		return nil, errors.New("failed to process registration")
	}

	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error hashing password: ", err)
		return nil, errors.New("failed to process registration")
	}

	// Create new user
	user := &models.User{
		Username:  request.Username,
		Password:  string(hashedPassword), // Password should be hashed in a real application
		VillageID: request.VillageID,
	}

	tx := s.userRepo.BeginTransaction()
	defer tx.Rollback()

	err = s.userRepo.CreateWithTx(tx, user)
	if err != nil {
		return nil, err
	}

	return &dtos.MessageResponse{
		Message: "User registered successfully",
	}, nil
}

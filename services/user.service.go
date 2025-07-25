package services

import (
	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"Apps-I_Desa_Backend/repositories"
	"errors"
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
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Create new user
	user := &models.User{
		Username: request.Username,
		Password: request.Password, // Password should be hashed in a real application
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

package userservice

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"main.go/dto"
	"main.go/entity"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUserByID(userID uint) (entity.User, error)
}

type Service struct {
	repo Repository
}

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		Password:    string(hash),
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Role:        entity.UserRole,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error %w", err)

	}

	return dto.RegisterResponse{
		User: dto.UserInfo{
			ID:          createdUser.ID,
			PhoneNumber: createdUser.PhoneNumber,
			Name:        createdUser.Name,
		},
	}, nil

}

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
type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	repo Repository
	auth AuthGenerator
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{auth: authGenerator, repo: repo}
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

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected err %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("invalid password: %w", err)
	}
	return dto.LoginResponse{
		User: dto.UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}

type ProfileRequest struct {
	UserID uint
}
type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, err
	}

	return ProfileResponse{Name: user.Name}, nil
}

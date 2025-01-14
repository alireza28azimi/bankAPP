package userservice

import "main.go/dto"

type Service struct {
}

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	return dto.RegisterResponse{}, nil
}

package petstore

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type PetstoreService struct{}

// ...
func NewPetstoreService() PetstoreServiceServer {
	return &PetstoreService{}
}

// ...
func (s *PetstoreService) ListPets(context.Context, *ListPetsRequest) (*ListPetsResponse, error) {
	return &ListPetsResponse{}, nil
}

// ...
func (s *PetstoreService) GetPet(context.Context, *GetPetRequest) (*Pet, error) {
	return &Pet{}, nil
}

// ...
func (s *PetstoreService) CreatePet(context.Context, *CreatePetRequest) (*Pet, error) {
	return &Pet{}, nil
}

// ...
func (s *PetstoreService) UpdatePet(context.Context, *UpdatePetRequest) (*Pet, error) {
	return &Pet{}, nil
}

// ...
func (s *PetstoreService) DeletePet(context.Context, *DeletePetRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

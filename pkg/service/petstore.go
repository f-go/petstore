package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PetstoreService struct{}

// ...
func NewPetstoreService() PetstoreServiceServer {
	return &PetstoreService{}
}

func (p PetstoreService) ListPets(ctx context.Context, req *ListPetsRequest) (*ListPetsResponse, error) {
	pet1 := Pet{
		Id: 1,
		Name: "Emil",
		Status: PetStatus_AVAILABLE,
		PhotoUrls: []string{
			"https://pictures/1.jpg",
			"https://pictures/2.jpg",
		},
		Category: &PetCategory{Name: "cat"},
	}
	pet2 := Pet{
		Id: 2,
		Name: "Freddy",
		Status: PetStatus_AVAILABLE,
		PhotoUrls: []string{
			"https://pictures/fredy1.jpg",
			"https://pictures/freddy2.jpg",
		},
		Category: &PetCategory{Name: "cat"},
	}

	return &ListPetsResponse{Pets: []*Pet{&pet1, &pet2}}, nil
}

func (p PetstoreService) GetPet(ctx context.Context, req *GetPetRequest) (*Pet, error) {
	if req.Id != 1 {
		err := status.Errorf(codes.NotFound, "Pet with ID %d not found.", req.GetId())
		return nil, err
	}

	pet := Pet{
		Id: 1,
		Name: "Emil",
		Status: PetStatus_AVAILABLE,
		PhotoUrls: []string{
			"https://pictures/1.jpg",
			"https://pictures/2.jpg",
		},
		Category: &PetCategory{Name: "cat"},
	}

	return &pet, nil
}

func (p PetstoreService) CreatePet(ctx context.Context, req *CreatePetRequest) (*Pet, error) {
	pet := req.GetPet()
	pet.Id = 1;

	return pet, nil
}

func (p PetstoreService) UpdatePet(ctx context.Context, req *UpdatePetRequest) (*Pet, error) {
	pet := req.GetPet()
	return pet, nil
}

func (p PetstoreService) DeletePet(ctx context.Context, req *DeletePetRequest) (*emptypb.Empty, error) {
	return &empty.Empty{}, nil
}

package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/pet-api/model"
	"github.com/pet-api/repository"
)

type PetService interface {
	Get() (pets []model.Pet, err error)
	Create(pet model.Pet) (createdPet model.Pet, err error)
	GetById(id string) (pet model.Pet, err error)
	Update(id string, pet model.Pet) (updatedPet model.Pet, err error)
	Delete(id string) (err error)
}

type petService struct {
	petRepository repository.PetRepository
}

// NewPetService creates a new PetService
// Constructs type PetService, depends on []
func NewPetService(petRepository repository.PetRepository) PetService {
	return &petService{
		petRepository: petRepository,
	}
}

func (p *petService) Get() (pets []model.Pet, err error) {
	if pets, err = p.petRepository.Get(); err != nil {
		return
	}

	return pets, nil
}

func (p *petService) GetById(id string) (pet model.Pet, err error) {
	var pets []model.Pet
	if pets, err = p.petRepository.Get(); err != nil {
		return
	}

	for _, v := range pets {
		if v.Id == id {
			pet = v
			return
		}
	}

	return pet, nil
}

func (p *petService) Create(pet model.Pet) (createdPet model.Pet, err error) {
	pet.Id = (uuid.New()).String()
	pet.CreatedAt = time.Now().Format(time.RFC3339)

	var pets []model.Pet
	if pets, err = p.petRepository.Get(); err != nil {
		return
	}

	pets = append(pets, pet)

	if err = p.petRepository.Create(pets); err != nil {
		return
	}

	createdPet = pet

	return createdPet, nil
}

func (p *petService) Update(id string, pet model.Pet) (updatedPet model.Pet, err error) {
	var pets []model.Pet
	if pets, err = p.petRepository.Get(); err != nil {
		return
	}

	petIndex := contains(pets, id)

	pets[petIndex].UpdatedAt = time.Now().Format(time.RFC3339)

	if pet.Age != 0 {
		pets[petIndex].Age = pet.Age
	}

	if pet.Breed != "" {
		pets[petIndex].Breed = pet.Breed
	}

	if pet.Name != "" {
		pets[petIndex].Name = pet.Name
	}

	if pet.Sex != "" {
		pets[petIndex].Sex = pet.Sex
	}

	if pet.Specie != "" {
		pets[petIndex].Specie = pet.Specie
	}

	if err = p.petRepository.Create(pets); err != nil {
		return
	}

	updatedPet = pets[petIndex]

	return updatedPet, nil
}

func (p *petService) Delete(id string) (err error) {
	var pets []model.Pet
	if pets, err = p.petRepository.Get(); err != nil {
		return
	}

	var updatedPets []model.Pet
	for i, v := range pets {
		if v.Id == id {
			updatedPets = remove(pets, i)
			if err = p.petRepository.Create(updatedPets); err != nil {
				return
			}
			return
		}
	}

	return nil
}

func remove(slice []model.Pet, index int) []model.Pet {
	return append(slice[:index], slice[index+1:]...)
}

func contains(slice []model.Pet, id string) int {
	for idx, v := range slice {
		if v.Id == id {
			return idx
		}
	}
	return -1
}

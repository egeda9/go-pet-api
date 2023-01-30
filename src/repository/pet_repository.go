package repository

import (
	"encoding/json"

	"github.com/pet-api/lib"
	"github.com/pet-api/model"
)

type PetRepository interface {
	Get() (pets []model.Pet, err error)
	Create(pets []model.Pet) (err error)
}

type petRepository struct {
	storage lib.Storage
}

// NewPetRepository creates a new PetRepository
// Constructs type PetRepository, depends on []
func NewPetRepository(storage lib.Storage) PetRepository {
	return &petRepository{
		storage: storage,
	}
}

func (p *petRepository) Get() (pets []model.Pet, err error) {
	data, err := p.storage.ReadJSON()

	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &pets); err != nil {
		return
	}

	return
}

func (p *petRepository) Create(pets []model.Pet) (err error) {
	if err = p.storage.WriteJSON(pets); err != nil {
		return
	}

	return nil
}

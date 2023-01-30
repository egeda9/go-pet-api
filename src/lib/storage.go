package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pet-api/model"
)

type Storage interface {
	ReadJSON() ([]byte, error)
	WriteJSON(pets []model.Pet) (err error)
}

type storage struct {
	env Env
}

func NewStorage(env Env) Storage {
	return &storage{
		env: env,
	}
}

func (s *storage) ReadJSON() (content []byte, err error) {
	if content, err = ioutil.ReadFile(s.env.StorageFilePath); err != nil {
		return
	}

	return
}

func (s *storage) WriteJSON(pets []model.Pet) (err error) {
	file, err := json.MarshalIndent(pets, "", " ")

	if err != nil {
		log.Println(fmt.Sprintf("Error when opening storage file %s", err.Error()), err)
		return err
	}

	err = ioutil.WriteFile(s.env.StorageFilePath, file, 0644)

	if err != nil {
		log.Println(fmt.Sprintf("Error when writing to the storage file %s", err.Error()), err)
		return err
	}

	return nil
}

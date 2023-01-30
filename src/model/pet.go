package model

import "fmt"

type Pet struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	Specie    SpecieType `json:"specie"`
	Sex       SexType    `json:"sex"`
	Breed     string     `json:"breed"`
	Age       float64    `json:"age"`
}

type SexType string

const (
	Male   SexType = "Male"
	Female SexType = "Female"
)

type SpecieType string

const (
	Dog SpecieType = "Dog"
	Cat SpecieType = "Cat"
)

func SpecieTypes() []string {
	species := []string{}
	species = append(species, string(Dog))
	species = append(species, string(Cat))
	return species
}

func SexTypes() []string {
	sexes := []string{}
	sexes = append(sexes, string(Male))
	sexes = append(sexes, string(Female))
	return sexes
}

type CreatePet struct {
	Name   string  `json:"name"`
	Specie string  `json:"specie"`
	Sex    string  `json:"sex"`
	Breed  string  `json:"breed"`
	Age    float64 `json:"age"`
}

func (s SexType) String() string {
	return string(s)
}

func ParseSexType(s string) (st SexType, err error) {
	sexTypes := map[SexType]struct{}{
		Male:   {},
		Female: {},
	}

	sex := SexType(s)
	_, ok := sexTypes[sex]
	if !ok {
		return st, fmt.Errorf(`cannot parse:[%s] as sex type`, s)
	}
	return sex, nil
}

func (s SpecieType) String() string {
	return string(s)
}

func ParseSpecieType(s string) (st SpecieType, err error) {
	speciesType := map[SpecieType]struct{}{
		Dog: {},
		Cat: {},
	}

	specie := SpecieType(s)
	_, ok := speciesType[specie]
	if !ok {
		return st, fmt.Errorf(`cannot parse:[%s] as specie type`, s)
	}
	return specie, nil
}

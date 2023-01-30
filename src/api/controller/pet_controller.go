package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pet-api/lib"
	"github.com/pet-api/model"
	"github.com/pet-api/service"
)

// PetController data type
type PetController struct {
	petService service.PetService
	util       lib.Util
}

// NewPetController creates new pet controller
// Constructs type NewPetController, depends on []
func NewPetController(petService service.PetService, util lib.Util) PetController {
	return PetController{
		petService: petService,
		util:       util,
	}
}

// @Summary      Find
// @Description  get pets
// @Tags         pet
// @Accept       json
// @Produce      json
// @Success      200
// @NoContent    204
// @NotFound     404
// @InternalServerError     500
// @Router       /pet [get]
func (p PetController) FindPet(ctx *gin.Context) {
	response, err := p.petService.Get()

	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error getting pets %s", err.Error()), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}

// @Summary      Find by id
// @Description  get by id. "id" property
// @Tags         pet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Unique Id"
// @Success      200
// @NoContent    204
// @NotFound     404
// @InternalServerError     500
// @Router       /pet/{id} [get]
func (p PetController) FindPetById(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := p.petService.GetById(id)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error getting a pet by id: %s", id), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if response.Id != "" {
		ctx.JSON(http.StatusOK, gin.H{"data": response})
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{"data": response})
	}
}

// @Summary      Create a new pet
// @Description  post a new pet
// @Tags         pet
// @Accept       json
// @Produce      json
// @Param        data body      model.CreatePet  true  "Pet".
// @Success      200  {json} model.Pet
// @BadRequest   400 {string} Bad request
// @NotFound     404 {string} Invalid route
// @InternalServerError     500 {string} Error
// @Router       /pet [post]
func (p PetController) PostPet(ctx *gin.Context) {
	var pet model.CreatePet
	if err := ctx.ShouldBindJSON(&pet); err != nil {
		log.Println(fmt.Sprintf("ERROR: Error getting the request body %s", err.Error()), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := p.ValidateRequestBody(pet)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error getting the request body %s", err.Error()), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPet, err := p.petService.Create(input)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error creating the pet %s", err.Error()), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newPet})
}

// @Summary      Update a pet
// @Description  update a pet
// @Tags         pet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "id"
// @Param        data body      model.CreatePet  true  "Pet".
// @Success      200  {json} model.Pet
// @BadRequest   400 {string} Bad request
// @NotFound     404 {string} Invalid route
// @InternalServerError     500 {string} Error
// @Router       /pet/{id} [put]
func (p PetController) UpdatePet(ctx *gin.Context) {
	id := ctx.Param("id")

	var pet model.CreatePet
	if err := ctx.ShouldBindJSON(&pet); err != nil {
		log.Println(fmt.Sprintf("ERROR: Error getting the request body %s", err.Error()), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := p.ValidateRequestBody(pet)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error getting the request body %s", err.Error()), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPet, err := p.petService.Update(id, input)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error updating the pet '%s' %s", id, err.Error()), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedPet})
}

// @Summary      Delete a pet
// @Description  Delete a pet
// @Tags         pet
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "id"
// @Success      200
// @BadRequest   400 {string} Bad request
// @NotFound     404 {string} Invalid route
// @InternalServerError     500 {string} Error
// @Router       /pet/{id} [delete]
func (p *PetController) DeletePet(ctx *gin.Context) {
	id := ctx.Param("id")

	err := p.petService.Delete(id)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error getting a pet by id: %s", id), err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "OK"})
}

func (p *PetController) ValidateRequestBody(pet model.CreatePet) (output model.Pet, err error) {
	if !p.util.Contains(model.SexTypes(), pet.Sex) {
		err = fmt.Errorf("ERROR: Invalid Sex type in request body '%s'", pet.Sex)
		log.Println(err.Error(), err)
		return
	}

	if !p.util.Contains(model.SpecieTypes(), pet.Specie) {
		err = fmt.Errorf("ERROR: Invalid Specie type in request body '%s'", pet.Specie)
		log.Println(err.Error(), err)
		return
	}

	specieType, err := model.ParseSpecieType(pet.Specie)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error parsing Specie type in request body '%s'", err.Error()), err)
		return
	}

	sexType, err := model.ParseSexType(pet.Sex)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR: Error parsing Sex type in request body '%s'", err.Error()), err)
		return
	}

	output = model.Pet{
		Name:   pet.Name,
		Specie: specieType,
		Sex:    sexType,
		Breed:  pet.Breed,
		Age:    pet.Age,
	}

	return
}

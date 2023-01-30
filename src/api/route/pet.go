package route

import (
	"fmt"

	"github.com/pet-api/api/controller"
	"github.com/pet-api/lib"
)

// HelloWorldRoutes struct
type PetRoutes struct {
	handler       lib.RequestHandler
	petController controller.PetController
}

// Setup HelloWorld routes
func (p PetRoutes) Setup() {
	fmt.Println("------------------- Setting up Pet routes -------------------")
	api := p.handler.Gin.Group("/api")
	{
		api.GET("/pet", p.petController.FindPet)
		api.GET("/pet/:id", p.petController.FindPetById)
		api.POST("/pet", p.petController.PostPet)
		api.PUT("/pet/:id", p.petController.UpdatePet)
		api.DELETE("/pet/:id", p.petController.DeletePet)
	}
}

// NewPetRoutes creates new Pet controller
// Constructs type NewPetRoutes, depends on []
func NewPetRoutes(
	handler lib.RequestHandler,
	petController controller.PetController,
) PetRoutes {
	return PetRoutes{
		handler:       handler,
		petController: petController,
	}
}

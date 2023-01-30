package route

import (
	"fmt"

	"github.com/pet-api/api/controller"
	"github.com/pet-api/lib"
)

// HelloWorldRoutes struct
type HelloWorldRoutes struct {
	handler              lib.RequestHandler
	helloWorldController controller.HelloWorldController
}

// Setup HelloWorld routes
func (h HelloWorldRoutes) Setup() {
	fmt.Println("------------------- Setting up HelloWorld routes -------------------")
	api := h.handler.Gin.Group("/api")
	{
		api.GET("/helloworld/:id", h.helloWorldController.FindById)
	}
}

// NewHelloWorldRoutes creates new HelloWorld controller
// Constructs type HelloWorldRoutes, depends on []
func NewHelloWorldRoutes(
	handler lib.RequestHandler,
	helloWorldController controller.HelloWorldController,
) HelloWorldRoutes {
	return HelloWorldRoutes{
		handler:              handler,
		helloWorldController: helloWorldController,
	}
}

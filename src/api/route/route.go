package route

import "go.uber.org/fx"

// Module exports dependency to container. Options converts a collection of Options into a single Option.
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewHelloWorldRoutes),
	fx.Provide(NewPetRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	helloWorldRoutes HelloWorldRoutes,
	petRoutes PetRoutes,
) Routes {
	return Routes{
		helloWorldRoutes,
		petRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

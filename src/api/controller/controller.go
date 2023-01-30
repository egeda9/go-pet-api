package controller

import (
	"go.uber.org/fx"
)

// Module exported for initializing application. Options converts a collection of Options into a single Option.
var Module = fx.Options(
	fx.Provide(NewHelloWorldController),
	fx.Provide(NewPetController),
)

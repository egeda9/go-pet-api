package lib

import (
	"go.uber.org/fx"
)

// Module exports dependency. Options converts a collection of Options into a single Option.
var Module = fx.Options(
	fx.Provide(NewEnv),
	fx.Provide(NewRequestHandler),
	fx.Provide(NewStorage),
	fx.Provide(NewUtil),
)

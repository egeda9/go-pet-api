package middleware

import "go.uber.org/fx"

// Module Middleware exported. Options converts a collection of Options into a single Option.
var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewMiddleware),
)

// IMiddleware middleware interface
type IMiddleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []IMiddleware

// NewMiddlewares creates new middlewares
// Register the middleware that should be applied directly (globally)
// Constructs type Middlewares, depends on CorsMiddleware
func NewMiddleware(corsMiddleware CorsMiddleware) Middlewares {
	return Middlewares{
		corsMiddleware,
	}
}

// Setup sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}

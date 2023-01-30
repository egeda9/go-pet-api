package middleware

import (
	"fmt"

	"github.com/pet-api/lib"
	cors "github.com/rs/cors/wrapper/gin"
)

// CorsMiddleware middleware for cors
type CorsMiddleware struct {
	handler lib.RequestHandler
	env     lib.Env
}

// NewCorsMiddleware creates new cors middleware
// Constructs type CorsMiddleware, depends on lib.RequestHandler, lib.Logger, lib.Env
func NewCorsMiddleware(handler lib.RequestHandler, env lib.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		env:     env,
	}
}

// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {
	fmt.Println("Setting up cors middleware")

	debug := m.env.Environment == "local" || m.env.Environment == "Development"
	m.handler.Gin.Use(cors.New(cors.Options{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		Debug:            debug,
	}))
}

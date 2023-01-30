package lib

import (
	"github.com/gin-gonic/gin"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
// Constructs type RequestHandler, depends on Logger
func NewRequestHandler() RequestHandler {
	engine := gin.New()
	return RequestHandler{Gin: engine}
}

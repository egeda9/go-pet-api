package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorldController data type
type HelloWorldController struct {
	// claimService services.ClaimService
	// logger       logger.Logger
	// env          lib.Env
	// util         lib.Util
	// validator    lib.Validator
}

// HelloWorldController creates new hello world controller
// Constructs type HelloWorldController, depends on []
func NewHelloWorldController(
// claimService services.ClaimService,
// logger logger.Logger,
// env lib.Env,
// util lib.Util,
// validator lib.Validator
) HelloWorldController {
	return HelloWorldController{
		// claimService: claimService,
		// logger:       logger,
		// env:          env,
		// util:         util,
		// validator:    validator,
	}
}

// @Summary      Find by id
// @Description  get by id. "id" property
// @Tags         hello_world
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Unique Id"
// @Success      200
// @NoContent    204
// @NotFound     404
// @InternalServerError     500
// @Router       /helloworld/{id} [get]
func (h HelloWorldController) FindById(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

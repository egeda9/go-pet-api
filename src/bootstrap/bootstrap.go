package bootstrap

import (
	"context"
	"fmt"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pet-api/api/controller"
	"github.com/pet-api/api/middleware"
	"github.com/pet-api/api/route"
	_ "github.com/pet-api/docs"
	"github.com/pet-api/lib"
	"github.com/pet-api/repository"
	"github.com/pet-api/service"

	"go.uber.org/fx"
)

// Module exported for initializing application. Options converts a collection of Options into a single Option.
var Module = fx.Options(
	controller.Module,
	route.Module,
	lib.Module,
	service.Module,
	middleware.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes route.Routes,
	env lib.Env,
	middleware middleware.Middlewares,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("------------------- Starting Application -------------------------")
			fmt.Println("------------------------------------------------------------------")
			fmt.Println("------- Pets API -------")
			fmt.Println("------------------------------------------------------------------")

			go func() {
				middleware.Setup()
				routes.Setup()

				fmt.Printf(fmt.Sprintf("%s%s", "Current environment: ", env.Environment))

				if env.Environment == "Development" || env.Environment == "local" {
					handler.Gin.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
				}

				handler.Gin.Run()
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Stopping Application")
			return nil
		},
	})
}

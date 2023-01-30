package main

import (
	"github.com/joho/godotenv"
	"github.com/pet-api/bootstrap"
	"go.uber.org/fx"
)

// @title           Pets API
// @version         1.0
// @description     Pets API

// @BasePath  /api
func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}

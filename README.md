# PETS API (Go)

## REST API Microservice

This is the Pets REST API microservice documentation which enables the administration of CRUD operations for pets data management.

It uses [Golang](https://go.dev) as base language and a [docker](https://www.docker.com) container-based deployment scheme

---

## Continuous Integration process

[Continuous Integration process (Branching patterns)](https://martinfowler.com/articles/branching-patterns.html)

## Golang Clean architecture

The solution architecture we can divide the code in 5 main layers:

- Model: Is a set of data structures.
- Service: Contains application specific business rules. It encapsulates and implements all of the use cases of the system.
- Controller: Is a set of adapters that convert data from the format most convenient for the use cases and entities.
- Lib: Is generally composed of frameworks and tools.
- Repository: Contains storage specific code and methods to operate on the data to and from the data storage

In Clean Architecture, each layer of the application (use case, data service and domain model) only depends on interface of other layers instead of concrete types. Dependency Injection is one of the SOLID principles, a rule about the constraint between modules that abstraction should not depend on details. Clean Architecture uses this rule to keep the dependency direction from outside to inside.

### Fx as a dependency injection application framework.

Fx uses the constructor injection pattern: https://pkg.go.dev/go.uber.org/fx#section-readme

#### Converting a Service into an Fx Application

Install the library via 

```
go get go.uber.org/fx
```

Next up, boot it up in the main.go file

```go
func main() {
	godotenv.Load()

	logger := lib.NewLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
```

Extracting providers

```go
// Module exports services present. Options converts a collection of Options into a single Option.
var Module = fx.Options(
	fx.Provide(NewClaimService),
)
```

Provider definition

```
fx.Provide (func Provide(constructors ...interface{}) Option)
```

### Prerequisites to run it locally

    - Docker
    - Visual Studio Code
    - Golang    

---

## Quick Start

After git clone, execute commands below in root directory:

### REST-API

```
+---api                                   //api folder
|   +---controllers                       //controllers folder
|       \---hello_world_controller.go     //-- hellow world controller 
|       \---pet_controller.go             //-- pets controller
|       \---controllers.go                //-- controllers module definition
|   +---middlewares                       //middlewares folder
|       \---cors.go                       //-- cors definitions
|       \---middlewares.go                //-- middlewares module definition
|   +---routes                            //routes folder
|       \---pet.go                        //-- pet routes definition
|       \---routes.go                     //-- routes module definition
+---bootstrap                             //bootstrap folder
|   \---bootstrap.go                      //-- module definition for app initialization
+---docs                                  //swagger docs folder (automatic generated via: swag init)
|   \---docs.go                           //-- golang swagger definition
|   \---swagger.json                      //-- json swagger definition
|   \---swagger.yaml                      //-- yaml swagger definition
+---lib                                   //lib folder
|   \---storage.go                        //-- in-memory storage implementation
|   \---env.go                            //-- env util implementation
|   \---lib.go                            //-- lib module definition
|   \---request_handler.go                //-- git request handler util implementation
|   \---util.go                           //-- util implementation
+---models                                //models folder
|   \---pet.go                            //-- pet model
+---services                              //services folder
|   \---pet_service.go                    //-- pet business implementation
|   \---service.go                        //-- service module definition
+---repository                            //repositories folder
|   \---pet_repository.go                 //-- pet repository implementation
|   \---repositories.go                   //-- repositories module definition
\---local.env                             //environment variables
\---go.mod                                //dependency management
\---go.sum                                //direct and indirect dependencies
\---main.go                               //startup
```

### Install

```
go mod download
swag init // To update the swagger files definitions
```

### Run the app

```
go build main.go
go run main.go
```

### Environment variables in local.env

```
ENV=development
STORAGE_FILE_PATH=./pets.json
```

### [infra](infra/) folder and env settings

- Deployment definition (per environment)
- Environment variables definition (per environment)

---

### Dependency modules representation

```
+---main                                   
|   +---bootstrap                       
|       +---controllers
|       	\---NewHelloWorldController
|       	\---NewPetController
|       +---routes
|       	\---NewPetRoutes
|       	\---NewRoutes
|       	\---NewHelloWorldRoutes
|       +---lib
|       	\---NewRequestHandler
|       	\---NewEnv
|       	\---NewHttpHandler
|       	\---NewUtil
|       +---middlewares
|       	\---NewCorsMiddleware
|       	\---NewMiddlewares
|       +---services
|         \---NewPetService  
```

---

### Accessing the API (From a local docker deployment)

Local Swagger Endpoint [Swagger](http://localhost:8080/swagger/index.html).

### Open API definition (rest-api)

The REST API app is described below.

[Open API Definition](./src/docs/swagger.json)

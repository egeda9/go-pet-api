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

### Docker deployment

#### Building the image


```
DOCKER_BUILDKIT=1 docker build -t go-pets-api -f Dockerfile .
```

This will build the image from the Dockerfile in the **/src** directory—the period at the end—and tag it with go-pets-api:latest, with go-pets-api being the image name and latest being the tag name.

#### Create and run the container

```
docker run -d -p 7991:80 --name go-pets-api go-pets-api:latest
```

- **-d**: This is short for detach and means that the Docker container will run in the background.
- **-p 7991:80**: This publishes the port 80 of the container as the port 7991 on the host. Thanks to that, the API will be available on the host at http://localhost:7991/swagger.
- **–name go-pets-api**: The container will be available under the name go-pets-api. 
- **go-pets-api**: This is the name of the image we want to use to create the container.

### Open API definition

The REST API app is described below.

[Open API Definition](./src/docs/swagger.json)

### gcloud - cloud run deployment

#### Deploy to Cloud Run from source

- In the source code directory (/src), deploy from source.

```
gcloud run deploy --source . --update-env-vars ENV=Development,STORAGE_FILE_PATH=../build/data/pets.json
```

- Specify service name
- Enable **Artifact Registry API**
- Select region **us-east1**
- Allow unauthenticated invocations


# REST API

The REST API to the pets api is described below.

## Get list of Pets

### Request

`GET /api/pet`

    curl -X 'GET' \
    '{host}/api/pet' \
    -H 'accept: application/json'

### Response

``` json
{
  "data": [
    {
      "id": "a9dbb813-8ac8-4365-8acb-923878b7b437",
      "name": "Lupe",
      "createdAt": "2023-01-27T11:44:02.0268285-05:00",
      "updatedAt": "2023-01-27T16:43:17.101Z",
      "specie": "Dog",
      "sex": "Female",
      "breed": "Poodle",
      "age": 1
    },
    {
      "id": "df78a87b-2a21-4480-aa43-d90dd85ff8c9",
      "name": "Canela",
      "createdAt": "2023-01-27T11:44:31.5572417-05:00",
      "updatedAt": "2023-01-27T16:43:17.101Z",
      "specie": "Dog",
      "sex": "Female",
      "breed": "Labrador Retriever",
      "age": 1.9
    }
  ]
}
```

## Get a specific Pet

### Request

`GET /api/pet/:id`

```
  curl -X 'GET' \
  '{host}/api/pet/df78a87b-2a21-4480-aa43-d90dd85ff8c9' \
  -H 'accept: application/json'
```

### Response

``` json
{
  "data": {
    "id": "df78a87b-2a21-4480-aa43-d90dd85ff8c9",
    "name": "Canela",
    "createdAt": "2023-01-27T11:44:31.5572417-05:00",
    "updatedAt": "2023-01-27T16:43:17.101Z",
    "specie": "Dog",
    "sex": "Female",
    "breed": "Labrador Retriever",
    "age": 1.9
  }
}
```

## Create a new Pet

### Request

`POST /api/pet`

```
  curl -X 'POST' \
	'{host}/api/pet' \
	-H 'accept: application/json' \
	-H 'Content-Type: application/json' \
	-d '{
	"age": 5,
	"breed": "German Shepherd",
	"name": "Riley",
	"sex": "Male",
	"specie": "Dog"
}'
```

### Response

``` json
{
  "data": {
    "id": "d9d7fe36-2892-483b-848a-641ba9a9a600",
    "name": "Riley",
    "createdAt": "2023-02-01T16:31:46Z",
    "updatedAt": "",
    "specie": "Dog",
    "sex": "Male",
    "breed": "German Shepherd",
    "age": 5
  }
}
```

## Update a Pets's property

### Request

`PUT /api/pet/:id`

```
  curl -X 'PUT' \
  '{host}/api/pet/d9d7fe36-2892-483b-848a-641ba9a9a600' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "age": 6,
  "breed": "German Shepherd",
  "name": "Riley",
  "sex": "Male",
  "specie": "Dog"
  }'
```

### Response

``` json
{
  "data": {
    "id": "d9d7fe36-2892-483b-848a-641ba9a9a600",
    "name": "Riley",
    "createdAt": "2023-02-01T16:31:46Z",
    "updatedAt": "2023-02-01T16:38:43Z",
    "specie": "Dog",
    "sex": "Male",
    "breed": "German Shepherd",
    "age": 6
  }
}
```

## Delete a Pet

### Request

`DELETE /api/pet/:id`

```
  curl -X 'DELETE' \
  '{host}/api/pet/d9d7fe36-2892-483b-848a-641ba9a9a600' \
  -H 'accept: application/json'
```

### Response

``` json
{
  "data": "OK"
}
```
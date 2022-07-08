# DelGong BE Story

## Requirements
1. Go 1.15 >
2. Mongodb version 3.4

## How to start in development, staging
- 
# Project Layout
- **application** directory contains entity, repository, dto, interactor, and messaging for your application, each dir contains
  so it can be registered into dependecy injection container
- **config** directory contains configuration for your application and it must have default values for local development.
  it accepts value from env variable and will override default value.
- **di** contains all dependecy injection configuration, all component must be registered here so it can be used by other components
- **interfaces** contains interface to outside world
- **migration** contains migration files
- **tests/mock** contains mocked file 

## good
1. write code that easy to remove, not easy to maintain
2. code through interface (SOLID)
3. all business code lives in interactor directory
4. controller has only 5 job :
    - bind req to dto
    - dto validation
    - authentication and authorization
    - calling interactor
    - send response
5. interactor handles business logic
6. repo handles persistent store (either db, file, cache etc ...)
7. entity represent something unique
7. dto bind request into model (model is not unique)
6. keep it simple (KISS)


## how to generate mockgen
1. install mockgen by running ```go install github.com/golang/mock/mockgen```
2. make sure to write code through interface
3. run ```mockgen -source=app/hotel/repository.go -package=mock_repository -destination=tests/mock/app/hotel/mock_repository.go```

## how to generate swagger
1. install swaggo by run this command ```go get github.com/swaggo/swag/cmd/swag```
2. comment your controller action [doc comment api](https://github.com/swaggo/swag#declarative-comments-format)
3. run ```swag init -g app/interfaces/hotel/*.go```
4. run ```go build main.go```
5. run ```./main run```

## NOTES
1. mongo repo is hard to be mocked, because mongo api always returns concrete type not interface type

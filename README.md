# gol_rest_api_boilerplate

steps:

- create module
- - go mod init name_of_module
- install fiber
- - go get github.com/gofiber/fiber/v2
- create ./cmd/api/main.go
- - touch main.go
- setting up main.go
- install swagger
- - go install github.com/swaggo/swag/cmd/swag@latest
- - swag init
- install swagger fiber middleware
- - go get -u github.com/swaggo/fiber-swagger
- add swagger routes

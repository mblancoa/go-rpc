


```
go-rpc
|-- cmd
|   |-- infofile
|       |-- server.go			//main file to run rpc service
|
|-- rpc
|   |-- infofile.go         // rpc server implementation
|
|--internal
   |-- errors
   |   |-- error.go
   |
   |-- core
   |   |-- ports
   |       |-- repository.go			//persistence definition
   |   |-- domain
   |       |-- domain.go
   |   |-- configuration.go
   |   |-- service.go 					//Business logic implementation
   |
   |-- adapters
   |   |--mongodb
   |      |-- configuration.go
   |      |-- repository.go 			//implementation with mongodb persistence of interfaces in ports
   |      |-- mongodbrepository.go		//domain and interface definition to generate code to mongodb access
```
## Description
This project contains a rpc service which receives a request 
```json
{
  "type": "any",  //default value: "core"
  "version": "any",   //default value: "1.0.0"
  "hash": "any"   //default value: ""
}
``` 
The service looks for the file with name `${type}-&{version}.json`, persists its content in and calculates its hash with SHA256 algorithm. 
The response is like below having empty content if the calculated hash is different to the one in request. 
```json
{
  "type": "any",
  "version": "any",
  "hash": "any",
  "content": "any"
}
```

The rpc service is a simple example of protobuf generation and serialization. And the business logic of the project is implemented with hexagonal 
architecture. 

In The file core/service.go the service infoFileService is implemented. It uses a InfoFileRepository defined in ports/repository.go to persist 
witch is implemented in adapters/mongodb/repository.go.

To be configured through inversion of dependencies pattern, it's expected that every adapter (only one in this case) has its own configuration to 
create every interface implementation that is needed by components in core.

Since JSON is the file's structure, the property `content` contains the file's text but not its binary content. `content` allows any object 
structure but protobuf serialization is not so flexible with unknown object structures so when postman is used to test, `content` is not like the 
real file's content.

## Repositories generation
Installation
```
go install github.com/sunboyy/repogen@latest
```
Generation
```
make code-generation
```
## Mocks generation
Installation
```
go install github.com/vektra/mockery/v2@v2.40.1
```
Generation
```
make clean mocks
```

## Run project
Docker image
```
docker build -t go-rpc .
```
Start network with docker-compose
```
    docker-compose up
```
Testing with postman
```
url: grpc://localhost:50051

{
  "type": "core",
  "version": "1.0.0",
  "hash": ""
}
```
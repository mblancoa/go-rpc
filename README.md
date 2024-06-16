


```
go-rpc
|-- cmd
|   |-- infofile
|       |-- server.go			//main file to run rpc service
|-- rpc
|   |-- infofile.go         // rpc server implementation
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
   |
   |-- rpc
   |   |-- infofile.go				//rpc service server implementation
   |
```
## Description
This project contains a rpc service which receives a request with the properties type, version and hash. 
The service finds a file with name type-version.json and inserts it into a data base. 

The rpc service has been implemented using protobuf generation and serialization. The business logic of the project is implemented with hexagonal architecture.

Since it has been my first rpc service, I have not had enough time to research how to send specific errors. I think protobuf  serialization is no the best to 
serialize unknown objects so it is needed to deserialize the response to know the correct content of the file. In the other hand, making requests from postman 
is very easy using reflexion.

Every component in the project has its own unit test using mocks when it is needed. It could be improved with integrations tests and functional tests (godog
)
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
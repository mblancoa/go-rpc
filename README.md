


```
go-rpc
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
   |-- cmd
   |   |-- infofile
   |       |-- server.go			//main file to run rpc service
```
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

## Build RPC server
```
go build ./cmd/infofile/server.go
```

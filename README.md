# calculator

This is a simple calculator to demonstrate the use of the following tools:

* Golang
* Protocol Buffers
* gRPC
* Go modules

## What it does

The calculator server implements a method Calculate add, substract, multiply or divide two numbers.

### System requirements

To get information about system requirements, please visit the [Golang MinimumRequirements Wiki Page](https://github.com/golang/go/wiki/MinimumRequirements)

### Installation

To install Golang, please follow the instructions on the [Golang Getting Started Page](https://golang.org/doc/install)

This project supports Go modules. Therefor it is not necessary anymore to "go get" all dependencies.
You can run the program either by running the two .go files for server and client:

1. ```$ go run server/calculator_server.go``` (first time Go modules are allocated)
2. ```$ go run client/calculator_client.go <number><operator><number>```

Or you build executable files by running:

1. ```$ go build server/calculator_server.go```
2. ```$ go build client/calculator_client.go```

Afterwards you can run the client like:  
```./calculator_client <number><operator><number>``` on Unix  
or  
```calculator_client.exe <number><operator><number>``` on Windows
# Go gRPC Echo Service

This is a simple Go gRPC project that implements an Echo service. It allows you to send a message to the server, and the server will echo the same message back.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [References](#references)

## Prerequisites
Before running this project, make sure you have the following installed on your system:
- Go: [Download and Installation Instructions](https://golang.org/doc/install)
- Protocol Buffers (protoc): [Download and Installation Instructions](https://developers.google.com/protocol-buffers/docs/downloads)
- gRPC Go: Install using `go get -u google.golang.org/grpc`

## Usage
1. Clone the repository:
   ```bash
   git clone -b Ankush--projects https://github.com/zversal-ecom/z-Skeleton.git
   cd z-Skeleton/gRPC-Go/echo


## Command to initalize the Project

1. Run the following command to genrate go.mod file 

`go mod init <project-name>`
in our case it is :

`go mod init echo`

2. Now run the following command to import or download the required dependencies:

`go mod tidy`


## Command to generate gRPC code

`protoc --go_out=./Echo --go_opt=paths=source_relative --go-grpc_out=./Echo --go-grpc_opt=paths=source_relative Echo/echo.proto`

## Run the Server

1. cd server

`go run main.go`

2. create the excutalbe file of the server.go file by this command :

`go build -o bin/server.exe server/main.go`

cd bin

run the server.exe file 

## Run the Client 

1. cd client

`go run main.go -echo <Message>`

2. create the excutalbe file of the client.go file by this command :

`go build -o bin/client.exe client/main.go`

cd bin

open the `CMD` run the following 

`client.exe -echo <Message>`

You will recieve the same message from the Server

protoc --go_out=./Echo --go_opt=paths=source_relative --go-grpc_out=./Echo --go-grpc_opt=paths=source_relative Echo/echo.proto
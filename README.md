# gRPC Example

This is a simple gRPC example project in Go. It includes a server and a client that communicate using gRPC. The server implements a `GreeterService` that responds to `SayHello` requests with a greeting message.

Associated documentation: <https://www.notion.so/Setting-Up-a-gRPC-API-with-Golang-on-macOS-Using-Visual-Studio-Code-188e46412b2280aeb888e0ca151578a0>

## Project Structure

```text
.vscode/
    launch.json

api/
    proto/
        v1/
            helloworld.proto

cmd/
    client/
        main.go
    server/
        main.go

go.mod
go.sum

pkg/
    greeter/
        api/
            proto/
                v1/
                    helloworld_grpc.pb.go
                    helloworld.pb.go

tools.go
```

- **api/proto/v1/helloworld.proto**: The Protocol Buffers definition file for the `GreeterService`.
- **cmd/client/main.go**: The client application that sends a `SayHello` request to the server.
- **cmd/server/main.go**: The server application that implements the `GreeterService`.
- **pkg/greeter/api/proto/v1/helloworld_grpc.pb.go**: The gRPC code generated from the `.proto` file.
- **pkg/greeter/api/proto/v1/helloworld.pb.go**: The Protocol Buffers code generated from the `.proto` file.
- **tools.go**: A file to manage build tools.
- **go.mod**: The Go module file.
- **go.sum**: The Go module dependencies file.
- **.vscode/launch.json**: Configuration for debugging in Visual Studio Code.

## Prerequisites

- Go 1.23.5 or later
- Protocol Buffers compiler (`protoc`)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

## Setup

1. Install the Protocol Buffers compiler (`protoc`).
2. Install the `protoc-gen-go` and `protoc-gen-go-grpc` plugins:

    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.4
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
    ```

3. Generate the gRPC and Protocol Buffers code:

    ```sh
    protoc --go_out=. --go-grpc_out=. api/proto/v1/helloworld.proto
    ```

## Running the Server

To run the server, execute the following command:

```sh
go run cmd/server/main.go
```

The server will start and listen on port `50051`.

## Running the Client

To run the client, execute the following command:

```sh
go run cmd/client/main.go [name]
```

Replace `[name]` with the name you want to send in the `SayHello` request. If no name is provided, it defaults to "world".

## Debugging

This project includes a `launch.json` file for debugging in Visual Studio Code. You can use the "Debug Server" and "Debug Client" configurations to debug the server and client, respectively.

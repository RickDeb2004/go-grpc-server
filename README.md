# Basic Go gRPC Server and Client

This is a basic gRPC server and client written in Go. It implements simple RPC, server-side streaming RPC, client-side streaming RPC, and bidirectional streaming RPC functionalities.

## Setting up the Project

### 1. Create Project Directory

```bash
mkdir basic-go-grpc
cd basic-go-grpc
mkdir client server proto
```

### 2. Installing gRPC Go Plugin

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

export PATH="$PATH:$(go env GOPATH)/bin"
```

### 3. Initialize Go Module

```bash
go mod init github.com/your_username/basic-go-grpc

go mod tidy
```

## Usage

### 1. Create Proto File

Create the proto file with the required services and messages in the `proto` directory.

### 2. Generate .pb.go Files

Generate `.pb.go` files from the proto file. Depending on the path mentioned in your `greet.proto` file, run either of the following commands:

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

OR

```bash
protoc --go_out=. --go_opt=module=github.com/your_username/basic-go-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/your_username/basic-go-grpc proto/greet.proto
```

### 3. Create Server and Client

Create the `server` and `client` directories. Inside each directory, create `main.go` files with necessary controllers and services.

### 4. Running the Application

#### Install Dependencies

```bash
go mod tidy
```

#### Run the Server

```bash
go run server/main.go
```

#### Run the Client

```bash
go run client/main.go
```

## Contributing

Feel free to contribute to this project by creating issues or submitting pull requests.

## License

This project is licensed under the [MIT License](LICENSE).

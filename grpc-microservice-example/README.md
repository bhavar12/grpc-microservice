# Production Ready Go gRPC Server and Client Microservice.

This is a Microservice whish has gRPC server and client written in Go. It is based on the

We have implemented a simple gRPC server and client with the following functionality:
- unary RPC
- for testing purpose we have exposed Gin Rest API since grpc client will not support postmen or browser
- TLS Authentication
- Postgres DB support
- In Memory DB
- Error Handling
- Config support for DB details

# Setting up a gRPC-Go project
1. Create a new directory for your project and cd into it

```bash
mkdir grpc-microservice-example
cd grpc-microservice-example
mkdir client server proto
```

2. Initialize a Go module

```bash
go mod init grpc-microservice-example

```

3. Download protobuf dependancy 

```bash
go get google.golang.org/protobuf/proto
```

4. Install protoc and "protoc" command works on your terminal
   ```Download the setup based on your OS from this link```
   https://github.com/protocolbuffers/protobuf/releases

   ```And setup the $PATH env varaible till bin```
   ```Verifi the protoc cmd working fine or not from your terminal```

5. Installing the gRPC Go plugin

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

```

6. Create the proto file with the required services and messages in the proto directory

7. Generate .pb.go files from the proto file

depending on what path you mention in your greet.proto file, you will either run this - 


```bash
protoc --go_out=. --go-grpc_out=. proto/movie.proto
```
8. Create the server and client directories and create the main.go files with necessary controllers and services


# Running the application

1. In a configs/app.json provide your postgres database details.

2. Run the below cmd to resolve the dependency
```bash
go mod tidy
```
3. Run the server

```bash
go run server/main.go
```

4. Run the client

```bash
go run client/main.go
```

5. You can use Make file if you have make installed on your machine

```bash
make server
```

```bash
make client
```

# Testing the application

1. Use Postmen or Below curl to create movie

```bash
curl --location 'http://localhost:5050/movies' \
--header 'Content-Type: application/json' \
--data '{
    "title":"Animal",
    "genre":"Cartton"
}'
```
capture the resp and used return id in below curls

2. To Fetch movie details
```bash
curl --location 'http://localhost:5050/movies/{id}'
```

3. To Delete Movie

```bash
curl --location --request DELETE 'http://localhost:5050/movies/{id}'
```

4. To get all movies

```bash
curl --location --request GET 'http://localhost:5050/movies' \
--header 'Content-Type: application/json'
```

# svc-user

- [Installation](#installation)
- [Usage](#usage)
- [Documentation](#documentation)

## Installation

### Requirements

1. [Go](https://golang.org/doc/install) 1.18+
2. [PosgreSQL](https://www.postgresql.org/download/) server
3. [Docker](https://golang.org/doc/install) for database testing with [testcontainer](https://www.testcontainers.org/) 
4. [Buf](https://docs.buf.build/introduction) for generating the grpc stubs
5. [protoc-gen-go](#), [protoc-gen-go-grpc](%), [](),

### Setting up environment

For a starter, export env `env.example` to your os env / docker env

## Getting Started
## Usage

### Development
```
 go mod tidy
```

After all installed properly, start the development.

```
 go run cmd/server/main.go -port 9998
```

#### Make requests:
HTTP/1.1 POST API with curl
```
$ curl \
--header "Content-Type: application/json" \
--data '{"name": "John"}' \
http://localhost:9998/user.v1.UserService/Create
```
gRPC with grpccurl
```
$ grpcurl \
-protoset <(buf build -o -) -plaintext \
-d '{"name": "John"}' \
localhost:9998 user.v1.UserService/Create
```

reponse:
```
{"message": "OK"}
```
## Documentation
### Todo:
- [ ] Observability
  - [ ] Opentelemetry Tracer & metric SDK
  - [ ] otel middleware for gorm, grpc/http
  - [ ] Server Otel Collector (docker-compose)
  - [ ] Server Jaeger (docker-compose)
  - [ ] Server Prometheus (docker-compose)
  - [ ] Server Grafana (docker-compose)
  - [ ] Log agent & agregator ELK 
- [ ] grpc integration test
- [ ] Dockerfile
- [ ] Makefile for test, build etc.
- [ ] proto to openapi spec (swagger UI)
- [ ] moreee :D

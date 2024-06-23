# GRPC With TLS secured
Implementing Transport Layer Secured (TLS) on gRPC client and server communication

## Prequisite
``` bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

brew install make
```

## Instalation
``` bash
make protoc
go mod tidy

# self signed Certificate
make certificate self
```
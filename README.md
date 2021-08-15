# server-examples

## To initialize the project, run:

`go mod init path-to-project-repo`

e.g. `go mod init github.com/alinabylkova/server-examples` where `server-examples` is a project name.

## Install required packages:

Run `go get package_name`

## Define grpc service in proto files:

Follow [Uber protobuf style guide](https://github.com/uber/prototool/blob/dev/style/README.md) when writing protocol buffer specifications

Creating new Services / APIs
1. Create a new directory for the service
2. Create separate files for messages and API (e.g. myservice_api.proto and myservice_messages.proto)

Recommended directory/file name: `company_name/service_name/version/service_name_api.proto`

Example: `mycompany/myservice/v1beta1/myservice_api.proto`
         `mycompany/myservice/v1beta1/myservice_messages.proto`

## Generate protofiles:

Before generating protos, install `protoc` compiler and protoc plugins for golang: 

- protoc-gen-go: `go install google.golang.org/protobuf/cmd/protoc-gen-go`
- protoc-gen-go-grpc: `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`
- protoc-gen-grpc-gateway: `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway`
- protoc-gen-openapiv2: `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2`

The latter 2 protoc-gen-grpc-gateway and protoc-gen-openapiv2 are used for grpc-http 

To enable annotations import, copy and paste third party files to this repository from:

- https://github.com/grpc-ecosystem/grpc-gateway/tree/master/protoc-gen-openapiv2/options
- https://github.com/googleapis/googleapis/tree/master/google/api

Then run `./genrate_proto.sh` script to generate proto files.
protoc --proto_path=service-interfaces --go_out=generated --go_opt=paths=source_relative service-interfaces/myproject/users/v1beta1/users_messages.proto
protoc --proto_path=service-interfaces --go_out=generated --go-grpc_out=generated --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false service-interfaces/myproject/users/v1beta1/users_api.proto

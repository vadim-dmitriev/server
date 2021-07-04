generate:
	protoc --proto_path=api --go_out=pkg --go-grpc_out=pkg --go_opt=paths=source_relative api.proto
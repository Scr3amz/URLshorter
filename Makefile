PHONY: generate-structs
generate-structs:
	mkdir -p internal/
	protoc --go_out=internal/ --go_opt=paths=source_relative \
	--go-grpc_out=internal/ --go-grpc_opt=paths=source_relative \
	api/proto/service.proto
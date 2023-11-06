PHONY: generate-structs
generate-structs:
	mkdir -p internal/urlshorter
	protoc --go_out=internal/urlshorter --go_opt=paths=source_relative \
	--go-grpc_out=internal/urlshorter --go-grpc_opt=paths=source_relative \
	api/urlshorter/service.proto
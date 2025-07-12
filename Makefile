.PHONY: generate
generate:
	@echo "Generating proto files..."
	@mkdir -p generated/go
	@protoc --go_out=generated/go --go_opt=paths=source_relative \
	        --go-grpc_out=generated/go --go-grpc_opt=paths=source_relative \
	        --proto_path=api/grpc \
	        api/grpc/google/api/*.proto \
	        api/grpc/proto/*.proto
	@echo "Proto generation completed!"

gen_protos:
	protoc \
		--proto_path=proto \
		--go_out=internal/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/pb \
		--go-grpc_opt=paths=source_relative \
		--validate_out=lang=go,paths=source_relative:internal/pb proto/task.proto

gen_swagger:
	swag init -g cmd/gateway/main.go

run_gateway:
	go run cmd/gateway/main.go

run_tasker:
	go run cmd/task-service/main.go

domain_tests:
	go test ./internal/domain/test -v

.PHONY: run run_gateway run_tasker
tidy:
	go mod tidy && go mod vendor

fmt:
	go fmt ./...

gen:
	protoc -I ./proto --go_out ./pb --go_opt paths=source_relative --go-grpc_out ./pb --go-grpc_opt paths=source_relative ./proto/*.proto

clean:
	rm pb/*.go

test:
	go test -cover -race ./...

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

.PHONY: tidy fmt gen clean test server client
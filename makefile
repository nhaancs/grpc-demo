tidy:
	go mod tidy && go mod vendor

gen:
	protoc -I ./proto --go_out ./pb --go_opt paths=source_relative --go-grpc_out ./pb --go-grpc_opt paths=source_relative ./proto/*.proto

clean:
	rm pb/*.go

test:
	go test -cover -race ./...

run:
	go run main.go
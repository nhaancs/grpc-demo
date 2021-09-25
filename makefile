gen:
	protoc -I ./proto --go_out ./pb --go_opt paths=source_relative --go-grpc_out ./pb --go-grpc_opt paths=source_relative ./proto/processor_message.proto

clean:
	rm pb/*.go

run:
	go run main.go
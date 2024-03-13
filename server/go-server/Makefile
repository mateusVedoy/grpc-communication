generate-rpc:
	@protoc --proto_path=protobufs \
	--go_out=./ \
	--go-grpc_out=./ \
	./protobufs/message.proto

run:
	@go run ./src/main.go
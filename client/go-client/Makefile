generate-rpc:
	@protoc --proto_path=protobufs \
		--go_out=./ \
		--go_opt=Mmessage.proto=src/protobuf \
		./protobufs/message.proto \
		--go-grpc_out=./ \
		--go-grpc_opt=Mmessage.proto=src/protobuf \
  		./protobufs/message.proto 

run:
	@go run ./src/main.go

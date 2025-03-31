proto-server:
	protoc --go_out=. --go-grpc_out=. --proto_path=server server/chat.proto

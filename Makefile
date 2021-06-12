protoc:
	protoc \
		--go_out=$(GOPATH)/src \
		--go-grpc_out=$(GOPATH)/src \
		protobuf/common.proto protobuf/kvs.proto protobuf/queue.proto protobuf/stack.proto

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get google.golang.org/grpc 
	@$(MAKE) update-path

update-path:
	@export PATH="$PATH:$(shell go env GOPATH)/bin" && \
	echo "PATH updated to include: $(shell go env GOPATH)/bin"

gen-agg-server:
	@protoc -I proto proto/aggregator_v1/aggregator.proto \
	 --go_out=./pkg/gen/ --go_opt=paths=source_relative \
	 --go-grpc_out=./pkg/gen/ --go-grpc_opt=paths=source_relative

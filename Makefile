PROTO_FILES := $(shell find pb -name "*.proto")
PROTO_TARGETS := $(PROTO_FILES:.proto=.pb.go)

pb: $(PROTO_TARGETS)

%.pb.go: %.proto
	protoc $< --go_out=plugins=grpc:.
	go test ./$(dir $<)

build:
	go build ./src/server
	go build ./src/client

all: pb build
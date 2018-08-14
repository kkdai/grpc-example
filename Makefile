PROTO_FILES := $(shell find pb -name "*.proto")
PROTO_TARGETS := $(PROTO_FILES:.proto=.pb.go)

pb: $(PROTO_TARGETS)

%.pb.go: %.proto
	protoc $< --go_out=plugins=grpc:.
	go test ./$(dir $<)

src/%: pb src/$@
	go build -i $(GO_BUILD_FLAGS) ./$@

all: pb src/client src/server
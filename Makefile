PROTO_FILES := $(shell find pb -name "*.proto")
PROTO_TARGETS := $(PROTO_FILES:.proto=.pb.go)

pb: $(PROTO_TARGETS)

%.pb.go: %.proto
	protoc $< --go_out=plugins=grpc:.
	go test ./$(dir $<)

cmd/%: pb src/$@
	go build -i $(GO_BUILD_FLAGS) ./src/$@

# cmds: $(CMD_TARGETS)

debug:
	@echo $(CMD_TARGETS)

CMD_PKGS := $(wildcard src/*)
CMD_TARGETS := $(subst src/,,$(CMD_PKGS))

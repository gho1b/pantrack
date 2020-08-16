PROJECT_DIR=.
PROTO_FILES=$(shell find $(PROJECT_DIR)/protobuf -name *.proto)
PROTO_GEN=./proto-gen

.PHONY: proto
proto:
	rm -rf $(PROTO_GEN)
	mkdir $(PROTO_GEN)
	for file in $(PROTO_FILES); do \
		protoc -I$(PROJECT_DIR)/protobuf --micro_out=paths=source_relative:$(PROTO_GEN) --go_out=paths=source_relative:$(PROTO_GEN) $${file}; \
	done

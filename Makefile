PROJECT_DIR=.
PROTO_FILES=$(shell find $(PROJECT_DIR) -name *.proto)

.PHONY: proto
proto:
	rm -rf $(PROJECT_DIR)/github.com
	for file in $(PROTO_FILES); do \
		protoc -I$(PROJECT_DIR) --micro_out=$(PROJECT_DIR) --go_out=$(PROJECT_DIR) $${file}; \
	done

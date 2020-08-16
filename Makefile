PROJECT_DIR=.
PROTO_SRC=$(PROJECT_DIR)/proto/
PROTO_OUT=$(PROJECT_DIR)/protobuf/
PROTO_FILES=$(shell find $(PROTO_SRC) -name *.proto -type f|sed s/.[\/]proto[\/]//g)

.PHONY: proto
proto:
	rm -rf $(PROTO_OUT)
	mkdir $(PROTO_OUT)
	echo $(PROTO_FILES)
	for file in $(PROTO_FILES); do \
  		echo $${file};\
		protoc --proto_path=$(PROJECT_DIR)/proto --micro_out=paths=source_relative:$(PROTO_OUT) --go_out=paths=source_relative:$(PROTO_OUT) $${file}; \
	done

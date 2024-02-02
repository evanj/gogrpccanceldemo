BUILD_DIR:=build
PROTOC:=$(BUILD_DIR)/bin/protoc
PROTOC_GEN_GO:=$(BUILD_DIR)/protoc-gen-go
PROTOC_GEN_GO_GRPC:=$(BUILD_DIR)/protoc-gen-go-grpc

all: $(BUILD_DIR)/.gocheck_stamp

$(BUILD_DIR)/.gocheck_stamp: $(wildcard **/*.go) echopb/echo.pb.go | $(BUILD_DIR)
	go test ./...
	go vet ./...
	staticcheck --checks=all ./...
	go mod tidy
	touch $@

echopb/echo.pb.go: proto/echo.proto $(PROTOC) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC)
	$(PROTOC) --plugin=$(PROTOC_GEN_GO) --plugin=$(PROTOC_GEN_GO_GRPC) \
		--go_out=. --go_opt=Mproto/echo.proto=./echopb \
		--go-grpc_out=. --go-grpc_opt=Mproto/echo.proto=./echopb \
		$<

# download protoc to a temporary tools directory
$(PROTOC): $(BUILD_DIR)/getprotoc | $(BUILD_DIR)
	$(BUILD_DIR)/getprotoc --outputDir=$(BUILD_DIR)

$(BUILD_DIR)/getprotoc: | $(BUILD_DIR)
	GOBIN=$(realpath $(BUILD_DIR)) go install github.com/evanj/hacks/getprotoc@latest

# go install uses the version of protoc-gen-go specified by go.mod ... I think
$(PROTOC_GEN_GO): go.mod | $(BUILD_DIR)
	GOBIN=$(realpath $(BUILD_DIR)) go install google.golang.org/protobuf/cmd/protoc-gen-go

# manually specified version since we don't import this from code anywhere
# TODO: Import this from some tool so it gets updated with go get?
$(PROTOC_GEN_GO_GRPC): go.mod | $(BUILD_DIR)
	GOBIN=$(realpath $(BUILD_DIR)) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

$(BUILD_DIR):
	mkdir -p $@

clean:
	$(RM) -r $(BUILD_DIR)

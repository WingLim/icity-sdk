BUILD_DIR   := build
FRAMEWORK	:= icitysdk.xcframework
BUILD_FLAGS := -v

CGO_ENABLED := 0
GO111MODULE := on

LDFLAGS += -w -s -buildid=

GO_BIND = GO111MODULE=$(GO111MODULE) \
	gomobile bind $(BUILD_FLAGS) -ldflags '$(LDFLAGS)' -trimpath

.PHONY: ios

all: ios

ios:
	$(GO_BIND) -target=ios -o $(BUILD_DIR)/$(FRAMEWORK)

clean:
	rm -rf $(BUILD_DIR)
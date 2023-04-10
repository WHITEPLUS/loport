GO := $(shell which go)
APP := loport
ARM := ./bin/$(APP).arm64
AMD := ./bin/$(APP).amd64

.PHONY:	build

build:
	GOOS=darwin GOARCH=amd64 $(GO) build -gcflags "-N -l" -o $(ARM)
	GOOS=darwin GOARCH=arm64 $(GO) build -gcflags "-N -l" -o $(AMD)
	lipo -create -output ./bin/$(APP) $(ARM) $(AMD)

GO := $(shell which go)
BIN := loport
APP := $(shell echo "$${PWD/$$GOPATH\/src\/}")

.PHONY:	build vendor test

build: vendor
	$(GO) build -i -gcflags "-N -l" -o $(BIN) ./*.go


vendor:
ifeq ($(shell type dep 2> /dev/null),)
	go get -u github.com/golang/dep/...
endif
	dep ensure


test:
	$(GO) test -cover -v $$(go list ./... | grep -v $(APP)/vendor)

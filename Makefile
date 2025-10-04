build_dir = ./bin
binary_name = allowimports
wd = $(realpath .)

.PHONY: build
build:
	go build -v ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: allowimports
allowimports:
	go build -v -o ${build_dir}/${binary_name} ./cmd/allowimports

.PHONY: plugin
plugin:
	go build -buildmode=plugin -tags=plugin -o ${build_dir}/${binary_name}.so ./cmd/plugin

.PHONY: demo
demo: allowimports
	go vet -vettool=${build_dir}/${binary_name} -config=${wd}/test/allowimports/allow.yaml ./test/allowimports/

.PHONY: all
all: build test demo
build_dir = ./bin
entrypoint = ./cmd/allowimports/main.go
binary_name = allowimports
here = $(realpath .)

.PHONY: build
build:
	go build -v ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: allowimports
allowimports:
	go build -v -o ${build_dir}/${binary_name} ${entrypoint}

.PHONY: plugin
plugin:
	go build -buildmode=plugin -o bin/allowimports.so ./cmd/plugin

.PHONY: demo
demo: allowimports
	go vet -vettool=${build_dir}/${binary_name} -config=${here}/test/allowimports/allow.yaml ./test/allowimports/
.PHONY: install-protoc-linux install-protoc-mac install-protoc-gen-go regenerate

install-protoc-linux:
	apt install -y protobuf-compiler

install-protoc-mac:
	brew install protobuf

install-protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest

regenerate:
	@rm -f $(shell find generated -name "*pb.go")
	@./generate.sh

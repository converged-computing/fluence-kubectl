HERE ?= $(shell pwd)
LOCALBIN ?= $(shell pwd)/bin

.PHONY: all build

all: build

.PHONY: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

.PHONY: protoc
protoc: $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	
# You can use make protoc to download proto
.PHONY: proto
proto: protoc
	PATH=$(LOCALBIN):${PATH} protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/fluence/service-grpc/service.proto

build:
	GO111MODULE="on" go build cmd/kubectl-fluence.go

# Don't do this, bad practice :)
install: build
	sudo cp $(HERE)/kubectl-fluence /usr/local/bin
	sudo cp $(HERE)/kubectl_complete-fluence /usr/local/bin

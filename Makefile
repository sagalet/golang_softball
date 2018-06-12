
ifeq ($(CROSS_COMPILE), true) 
GO := $(GOROOT)/bin/gox
else
GO := $(GOROOT)/bin/go
endif

PROTOC := /home/davis_chen/bin/protoc

GOPATH := $(shell pwd)
OUT := out

PACKAGES_CONFIG_SRC := \
	src/packages/config/reader.go \
	src/packages/config/data.go

all: packages web

packages: packages/config

out/proto/src/content1:
	/bin/mkdir -p $@
	$(PROTOC) -I=./proto --go_out=$@ ./proto/message.proto

packages/config: $(PACKAGES_CONFIG_SRC)
	GOPATH=$(GOPATH) $(GO) install $@

web:
	GOPATH=$(GOPATH):$(GOPATH)/out/proto GOBIN=$(GOPATH)/out/bin $(GO) $(CROSS_PARAMETER) install -v $@

.PHONY: all

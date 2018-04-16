
ifeq ($(CROSS_COMPILE), true) 
GO := $(GOROOT)/bin/gox
else
GO := $(GOROOT)/bin/go
endif

GOPATH := $(shell pwd)
OUT := out

PACKAGES_CONFIG_SRC := \
	src/packages/config/reader.go \
	src/packages/config/data.go

all: packages web

packages: packages/config

packages/config: $(PACKAGES_CONFIG_SRC)
	GOPATH=$(GOPATH) $(GO) install $@

web:
	GOPATH=$(GOPATH) $(GO) $(CROSS_PARAMETER) build -o $(OUT)/$@ $@

.PHONY: all


GO := $(GOROOT)/bin/go
GOPATH := $(shell pwd)
OUT := out

all: web

web:
	GOPATH=$(GOPATH) $(GO) build -o $(OUT)/$@ $@

.PHONY: all

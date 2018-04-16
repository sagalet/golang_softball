
ifeq ($(CROSS_COMPILE), true) 
GO := $(GOROOT)/bin/gox
else
GO := $(GOROOT)/bin/go
endif

GOPATH := $(shell pwd)
OUT := out

all: web

%:
	GOPATH=$(GOPATH) $(GO) $(CROSS_PARAMETER) build -o $(OUT)/$@ $@

.PHONY: all

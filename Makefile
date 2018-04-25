
DEP_PATH := github.com/golang/dep/cmd/dep
SRC_PATH := github.com/wzshiming/pic2ascii/cmd/...
GOBIN    := $(GOPATH)/bin
PATH     := $(PATH):$(GOBIN)
GO       := /usr/bin/env go
DEP      := /usr/bin/env dep


install: ensure
	$(GO) get -v $(SRC_PATH)

ensure: dep
	$(DEP) ensure

dep:
	$(GO) get -v $(DEP_PATH)
.PHONY: build test clean vet fmt chkfmt

APP         = hasher
VERSION     = $(shell git describe --tags --abbrev=0)
GO          = go
GO_BUILD    = $(GO) build
GO_FORMAT   = $(GO) fmt
GOFMT       = gofmt
GO_LIST     = $(GO) list
GO_TEST     = $(GO) test -v
GO_TOOL     = $(GO) tool
GO_VET      = $(GO) vet
GO_DEP      = $(GO) mod
GOOS        = ""
GOARCH      = ""
GO_PKGROOT  = ./...
GO_PACKAGES = $(shell $(GO_LIST) $(GO_PKGROOT))
GO_LDFLAGS  = -ldflags '-X github.com/nao1215/hasherly/cmd/hasher.Version=${VERSION}'

build:  ## Build binary
	env GO111MODULE=on CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD) $(GO_LDFLAGS) -o $(APP) cmd/hasher/main.go

clean: ## Clean project
	-rm -rf $(APP) cover.out cover.html

test: ## Start test
	env GOOS=$(GOOS) $(GO_TEST) -cover $(GO_PKGROOT) -coverpkg=./... -coverprofile=cover.out
	$(GO_TOOL) cover -html=cover.out -o cover.html

changelog: ## Generate changelog
	ghch --format markdown > CHANGELOG.md

fmt: ## Format go source code 
	$(GO_FORMAT) $(GO_PKGROOT)

.DEFAULT_GOAL := help
help:  
	@grep -E '^[0-9a-zA-Z_-]+[[:blank:]]*:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;32m%-15s\033[0m %s\n", $$1, $$2}'

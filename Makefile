BIN_DIR := $(GOPATH)/bin
GOLINT  := $(BIN_DIR)/golint
PGG     := $(BIN_DIR)/protoc-gen-go
GBD     := $(BIN_DIR)/go-bindata
PKGS    := $(shell go list ./... | grep -v vendor)
BINARY  := $(BIN_DIR)/bin/eventmaster

VERSION := $(shell git describe --tags 2> /dev/null || echo "unreleased")
V_DIRTY := $(shell git describe --exact-match HEAD 2> /dev/null > /dev/null || echo "-dirty")
GIT     := $(shell git rev-parse --short HEAD)
DIRTY   := $(shell git diff-index --quiet HEAD 2> /dev/null > /dev/null || echo "-dirty")


$(BINARY): deps $(wildcard **/*.go) proto ui/ui.go templates/templates.go
	@go install -v -ldflags \
		"-X github.com/ContextLogic/eventmaster.Version=$(VERSION)$(V_DIRTY) \
		 -X github.com/ContextLogic/eventmaster.Git=$(GIT)$(DIRTY)" \
		github.com/ContextLogic/eventmaster/cmd/eventmaster

.PHONY: proto
proto: deps proto/eventmaster.pb.go

proto/eventmaster.pb.go: $(PGG) proto/eventmaster.proto
	protoc --plugin=${PGG} -I proto/ proto/eventmaster.proto --go_out=plugins=grpc:proto

.PHONY: test
test: deps proto/eventmaster.pb.go ui/ui.go templates/templates.go
	@go test ${PGKS}

.PHONY: lint
lint: deps $(GOLINT)
	@go vet .
	@golint -set_exit_status .

# TODO: golint and protoc-gen-go are fetched from master still; should pin them down.
$(GOLINT):
	go get -u github.com/golang/lint/golint

$(PGG):
	go get -u github.com/golang/protobuf/protoc-gen-go

$(GBD): vendor
	go install ./vendor/github.com/jteeuwen/go-bindata/go-bindata

.PHONY: run
run: $(BINARY)
	eventmaster -r

.PHONY: deps
deps: vendor
vendor: Gopkg.toml Gopkg.lock
	dep ensure

ui:
	@mkdir ui

ui/ui.go: $(GBD) $(wildcard static/ui/**/*) ui
	go-bindata -prefix="static/" -o ui/ui.go -pkg=ui static/ui/...

templates:
	@mkdir templates

templates/templates.go: $(GBD) $(wildcard static/templates/*) templates
	go-bindata -prefix="static/" -o templates/templates.go -pkg=templates static/templates/...

.PHONY: coverage
coverage: 
	@go test -coverprofile=/tmp/cover github.com/ContextLogic/eventmaster 
	@go tool cover -html=/tmp/cover -o coverage.html
	@rm /tmp/cover

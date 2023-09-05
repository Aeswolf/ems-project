HASGOLINT := $(shell which golangci-lint 2> /dev/null)
DIRS := $$(find . -type f -name '*.go' -not -path "*/vendor/*" -not -path "*/.history/*")


ifdef HASGOLINT
	GOLINT=golangci-lint
else
	GOLINT=bin/golangci-lint
endif

build:
	go build -o ./bin/ ./.

run: build
	./bin/equipment-database-management

test:
	go test -v --race ./...

fix:
	@for dir in $(DIRS); do \
		gofmt -s -w $$dir; \
		goimports -w $$dir; \
	done

.PHONY: fix
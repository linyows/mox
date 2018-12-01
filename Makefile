TEST ?= ./...

default: build

deps:
	go get -u golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/cover
	go get github.com/goreleaser/goreleaser

build:
	go build

test:
	go test $(TEST) $(TESTARGS)
	go test -race $(TEST) $(TESTARGS) -coverprofile=coverage.txt -covermode=atomic

lint:
	golint -set_exit_status $(TEST)

ci: deps test lint

dist:
	@test -z $(GITHUB_TOKEN) || goreleaser --rm-dist

clean:
	rm -rf coverage.txt
	git checkout go.*

.PHONY: default dist test deps

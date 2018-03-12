build:
	go build -o tmplinux main.go

clean:
	rm tmplinux

tmpbuild: build clean

deps:
	dep ensure

unit_test:
	go test -v ./...

integration_test:
	bash integration_test.sh

test: unit_test integration_test

lint:
	bash -c "! go list ./... | xargs -L1 golint 2>&1 | read"

fmt:
	bash -c "! gofmt -d . 2>&1 | read"

vet:
	go vet ./...

static: lint fmt vet

check: static test tmpbuild

correct:
	gofmt -w .

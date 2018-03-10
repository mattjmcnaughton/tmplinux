build:
	go build -o tmplinux main.go

tmpbuild: build
	rm tmplinux

deps:
	dep ensure

unit_test:
	go test ./...

integration_test:
	echo 'Still need to write intergration tests'

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

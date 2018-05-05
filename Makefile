generate_build_files:
	dep ensure && bazel run //:gazelle

build: generate_build_files
	bazel build tmplinux

clean:
	bazel clean

unit_test: generate_build_files
	bazel test //pkg/...

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

check: static test

correct:
	gofmt -w .

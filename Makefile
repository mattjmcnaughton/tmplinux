build:
	go build -o tmplinux main.go

test:
	go test

lint:
	bash -c "! gofmt -d . 2>&1 | read"

correct:
	gofmt -w .

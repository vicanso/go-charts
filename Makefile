export GO111MODULE = on

.PHONY: default test test-cover dev build

build:
	go build -tags netgo -o go-charts

release:
	go mod tidy

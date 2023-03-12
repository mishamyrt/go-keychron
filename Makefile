.PHONY: build
build:
	go build -o build/keychron-rgb-adapter

.PHONY: lint
lint:
	golangci-lint run
	revive -config revive.toml  ./...
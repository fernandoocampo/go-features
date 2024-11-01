.PHONY: test
test: ## Run test
	go test -race ./...

.PHONY: test-1-23
test-1-23: ## Run test
	go1.23.0 test -race -timeout 10s github.com/fernandoocampo/go-features/onetwentythree/...
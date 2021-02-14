
.DEFAULT_GOAL := help

.PHONY: test
test: ## run tests (using go1.16beta1 for now)
	go1.16beta1 test -v -race .

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

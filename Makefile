.DEFAULT_GOAL := help
.PHONY: up rebuild down ssh test test.coverage build logs docker-cleanup help lint vet fmt deps vendor

DOCKER_C := docker-compose
DOCKER_RUN := docker-compose run 
APP_NAME := app

OUTPUT_COVERAGE := pkg/tests/coverage/

up: ## Start docker container
	$(DOCKER_C) pull
	$(DOCKER_C) up -d

rebuild: ## Rebuild docker container
	$(DOCKER_C) pull
	$(DOCKER_C) up --build -d

down: ## Stop docker container
	$(DOCKER_C) down

ssh: ## Interactive access to container
	$(DOCKER_RUN) --entrypoint="/bin/sh" $(APP_NAME)

deps: ## Install dependencies
	$(DOCKER_RUN) --entrypoint="go mod download" $(APP_NAME)

vendor: ## Install vendor
	$(DOCKER_RUN) --entrypoint="go mod vendor" $(APP_NAME)

lint: ## Checks Code Style
	$(DOCKER_RUN) --entrypoint="./run-lint.sh" $(APP_NAME)

vet: ## Finds issues in code
	$(DOCKER_RUN) --entrypoint="go vet ./..." $(APP_NAME)

fmt: ## Applies standard formatting
	$(DOCKER_RUN) --entrypoint="go fmt ./..." $(APP_NAME)

doc: ## Show package documentation
	$(DOCKER_RUN) --entrypoint="go doc github.com/jmsilvadev/form3libs/apiclient" $(APP_NAME)

test: ## Run all available tests
	$(DOCKER_RUN) --entrypoint="./run-tests.sh" $(APP_NAME)

test.coverage: ## Check project test coverage
	$(DOCKER_RUN) --entrypoint="./run-tests.sh" $(APP_NAME)
	open $(OUTPUT_COVERAGE)coverage.html >&- 2>&- || \
	xdg-open $(OUTPUT_COVERAGE)coverage.html >&- 2>&- || \
	gnome-open $(OUTPUT_COVERAGE)coverage.html >&- 2>&-

build: ## Build docker image in daemon mode
	$(DOCKER_C) build

logs: ## Watch docker log files
	$(DOCKER_C) logs --tail 100 -f

help:
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all test lint compile build push

export LINTER_VERSION ?= 1.45.2

IMAGE        = $(REGISTRY)/book-management
DOCKERFILES  = $(shell cd build && find */ -name 'Dockerfile' -print)
MODULES      = $(shell cd module && ls -d */)
ODIR        := build/_output
GO_PACKAGES ?= $(shell go list ./... | grep -v 'mock' | grep -v 'config' | grep -v 'sandbox')
TMP_DIR     := $(shell mktemp -d)
UNAME       := $(shell uname)

export VERSION      ?= $(shell git show -q --format=%h)
export VAR_SERVICES ?= $(foreach path, $(DOCKERFILES), $(path:%/Dockerfile=%))
export REGISTRY     ?= yusufemon

# Default postgres migration settings
export POSTGRES_USER ?= postgres
export POSTGRES_PASS ?= postgres
export POSTGRES_HOST ?= localhost
export POSTGRES_PORT ?= 5432
export POSTGRES_QUERYSTRING ?= sslmode=disable

all: check test

tool-lint:
	@test -e ./bin/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v${LINTER_VERSION}

lint: tool-lint
	./bin/golangci-lint run -v --timeout 3m0s

bin:
	@mkdir -p bin

tool-migrate: bin
ifeq ($(UNAME), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else ifeq ($(UNAME), Darwin)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.darwin-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else
	@echo "Your OS is not supported."
endif

migrate:
	@cp module/db/migrations/*.sql $(TMP_DIR)
	@bin/migrate -source file://$(TMP_DIR) -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@$(POSTGRES_HOST):$(POSTGRES_PORT)/sku_management?$(POSTGRES_QUERYSTRING)" $(MIGRATE_ARGS)

test:
	@go test -race -v ${GO_PACKAGES}

tool-mgen:
	@go install github.com/golang/mock/mockgen@v1.6.0

tool-coverage:
	@go install github.com/boumenot/gocover-cobertura@v1.2.0
	@go install github.com/jstemmer/go-junit-report@v0.9.1

coverage: tool-coverage
	@go test -timeout 30s --tags=integration -v -race -cover -coverprofile=coverage.out ${GO_PACKAGES} 2>&1 | go-junit-report -set-exit-code > report.xml
	@go tool cover -func=coverage.out
	@gocover-cobertura < coverage.out > coverage.xml

cover:
	@go tool cover -html=coverage.out

$(ODIR):
	@mkdir -p $(ODIR)

compile: $(ODIR)
	@$(foreach svc, $(VAR_SERVICES), \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(ODIR)/$(svc) cmd/$(svc)/main.go;)

build:
	@$(foreach svc, $(VAR_SERVICES), \
		docker build -t $(IMAGE)-$(svc):$(VERSION) -f ./build/$(svc)/Dockerfile .;)

push:
	@$(foreach svc, $(VAR_SERVICES), \
		docker push $(IMAGE)-$(svc):$(VERSION);)

docker-up:
	@docker-compose -f dev/docker-compose.yml up -d

docker-down:
	@docker-compose -f dev/docker-compose.yml down

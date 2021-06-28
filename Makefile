EXE_FILE = ./bin/custom-webhooks
COMMIT_HASH_SHORT = $(shell git rev-parse --short HEAD)
PROJECT = custom-webhook-store-logs

all: test lint secure

scan:
	@scripts/scanner.sh

test:
	@scripts/coverage.sh

lint:
	@golangci-lint --verbose run

mocks:
	@scripts/gen-mocks.sh

secure:
	@gosec ./...

clean:
	@rm $(EXE_FILE)

doc:
	@scripts/swagger.sh

build_for_mac:
	@CGO_ENABLED=0 go build -ldflags "-X $(PROJECT)/pkg/constants.CommitHashShort=$(COMMIT_HASH_SHORT)" -o $(EXE_FILE) ./cmd

build:
	@CGO_ENABLED=0 GOOS=linux go build -ldflags "-X $(PROJECT)/pkg/constants.CommitHashShort=$(COMMIT_HASH_SHORT)" -o $(EXE_FILE) ./cmd

run:
	@scripts/run.sh

build_and_run: build
	$(EXE_FILE)

docker_build:
	@docker build -t custom-webhooks:$(COMMIT_HASH_SHORT) .

docker_run:
	@docker run -p ${container_port}:${port} custom-webhooks:$(COMMIT_HASH_SHORT)

# format all dependencies
fmt:
	@go fmt $(shell go list ./... | grep -v /vendor/ | grep -v mocks)
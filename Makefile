EXE_FILE = ./bin/vtp-apis
COMMIT_HASH_SHORT = $(shell git rev-parse --short HEAD)
PROJECT = vtp-apis


scan:
	@scripts/scanner.sh


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
	@docker build -t vtp-apis:$(COMMIT_HASH_SHORT) .

docker_run:
	@docker run -p ${container_port}:${port} vtp-apis:$(COMMIT_HASH_SHORT)

# format all dependencies
fmt:
	@go fmt $(shell go list ./... | grep -v /vendor/ | grep -v mocks)
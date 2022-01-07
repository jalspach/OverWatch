
### https://betterprogramming.pub/my-ultimate-makefile-for-golang-projects-fcc8ca20c9bb
GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=example
VERSION?=0.0.0
SERVICE_PORT?=3000
DOCKER_REGISTRY?= #if set it should finished by /
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)


## Hello:
hello: ## Hello shows that you are tallking to this Makefile
	echo "Hello...I am working"

##Vendor:
vendor: ## Download the dependencies and put them in the /vendor directory.
	$(GOCMD) mod vendor

## Build:
build: ## Build for the current system
	go build -o bin/main main.go

## Run:
run: ## Run the app
	go run main.go

## Clean:
clean: ## Remove build related files
	rm -fr ./bin
	rm -fr ./out
	rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov yamllint-checkstyle.xml


## Build-all:
build-all: build-rpi build-linux-64bit build-windows-64bit build-mac ## Builds all of the builtin packages

## Build-rpi:
build-rpi: build-linux-arm5 ## Build for SCOE Pi's (alias for arm5)

## Build-linux-arm5:
build-linux-arm5: ## Aliased to build-pi
	GOOS=linux GOARCH=arm GOARM=5 go build 

## Build-linux-64bit:
build-linux-64bit: ## Build for 64bit Linux
	GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux app.go

## Build-windows-64bit:
build-windows-64bit: ## Build for 64bit Windows
	GOOS=darwin GOARCH=amd64 go build -o bin/app-amd64-darwin app.go

## Build-mac:
build-mac: ## Build for 64bit Mac
	GOOS=darwin GOARCH=amd64 go build

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)


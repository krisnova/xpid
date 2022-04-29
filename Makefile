# =========================================================================== #
#            MIT License Copyright (c) 2022 Kris Nóva <kris@nivenly.com>      #
#                                                                             #
#                 ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓                 #
#                 ┃   ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗   ┃                 #
#                 ┃   ████╗  ██║██╔═████╗██║   ██║██╔══██╗  ┃                 #
#                 ┃   ██╔██╗ ██║██║██╔██║██║   ██║███████║  ┃                 #
#                 ┃   ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║  ┃                 #
#                 ┃   ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║  ┃                 #
#                 ┃   ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝  ┃                 #
#                 ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛                 #
#                                                                             #
#                        This machine kills fascists.                         #
#                                                                             #
# =========================================================================== #

all: compile

version     =  1.0.5
target      =  xpid
org         =  kris-nova
authorname  =  Kris Nóva
authoremail =  kris@nivenly.com
license     =  MIT
year        =  2022
copyright   =  Copyright (c) $(year)
cstyle      =  Google

compile: ## Compile for the local architecture ⚙
	@echo "Compiling..."
	go build -ldflags "\
	-X 'github.com/$(org)/$(target).Version=$(version)' \
	-X 'github.com/$(org)/$(target).AuthorName=$(authorname)' \
	-X 'github.com/$(org)/$(target).AuthorEmail=$(authoremail)' \
	-X 'github.com/$(org)/$(target).Copyright=$(copyright)' \
	-X 'github.com/$(org)/$(target).License=$(license)' \
	-X 'github.com/$(org)/$(target).Name=$(target)'" \
	-o $(target) cmd/*.go

install: ## Install the program to /usr/bin 🎉
	@echo "Installing..."
	sudo cp $(target) /usr/bin/$(target)

test: clean compile install ## 🤓 Run go tests
	@echo "Testing..."
	go test -v ./...

clean: ## Clean your artifacts 🧼
	@echo "Cleaning..."
	rm -rf libxpid/build/*
	rm -rvf release/*

format: ## Format the code
	@echo "  ->  Formatting code"
	clang-format -i -style=$(cstyle) libxpid/include/*.h
	clang-format -i -style=$(cstyle) libxpid/src/*.h
	clang-format -i -style=$(cstyle) libxpid/src/*.c


.PHONY: libxpid
libxpid: libxpid-clean ## Compile and install libxpid
	@echo "Building libxpid..."
	mkdir -p libxpid/build
	cd libxpid && ./configure
	cd libxpid/build && make

.PHONY: aur
aur: ## Publish to AUR using my SSH key
	@echo "Publishing to AUR using Kris Nóva's key (if exists)..."
	cd aur && ./publish

.PHONY: libxpid-clean
libxpid-clean: ## Clean libxpid
	@echo "Clean libxpid..."
	rm -rf libxpid/build/*

.PHONY: libxpid-install
libxpid-install: ## Install libxpid
	@echo "Installing libxpid..."
	cd libxpid/build && make install

.PHONY: release
release: ## Make the binaries for a GitHub release 📦
	mkdir -p release
	GOOS="linux" GOARCH="amd64" go build -ldflags "-X 'github.com/$(org)/$(target).Version=$(version)'" -o release/$(target)-linux-amd64 cmd/*.go
	GOOS="linux" GOARCH="arm" go build -ldflags "-X 'github.com/$(org)/$(target).Version=$(version)'" -o release/$(target)-linux-arm cmd/*.go
	GOOS="linux" GOARCH="arm64" go build -ldflags "-X 'github.com/$(org)/$(target).Version=$(version)'" -o release/$(target)-linux-arm64 cmd/*.go
	GOOS="linux" GOARCH="386" go build -ldflags "-X 'github.com/$(org)/$(target).Version=$(version)'" -o release/$(target)-linux-386 cmd/*.go
	GOOS="darwin" GOARCH="amd64" go build -ldflags "-X 'github.com/$(org)/$(target).Version=$(version)'" -o release/$(target)-darwin-amd64 cmd/*.go

.PHONY: help
help:  ## 🤔 Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'

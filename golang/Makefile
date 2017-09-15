all: run

# This Makefile contains some convenience commands for deploying and publishing.

# For example, to build and run the docker container locally, just run:
# $ make

# or to publish the :latest version to the specified registry as :1.0.0, run:
# $ make publish version=1.0.0

app = docker-first-demo
binary_version = v0.0.2
version ?= latest
local_port = 8081
go_version ?= 1.8
git_hash = $(shell git rev-parse HEAD | cut -c 1-6)

binary:
	$(call blue, "# Building Golang Binary...")
	docker run --rm -it -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${go_version} sh -c 'go get && CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s -X main.version=${binary_version} -X main.gitHash=${git_hash} -X main.date=`date -u '+%Y-%m-%d_%I:%M:%S%p'`" -o ${app}'

image: binary
	$(call blue, "# Building Docker Image...")
	docker build --no-cache -t ${app}-golang:${version} .
	$(MAKE) clean

run: image
	$(call blue, "# Running Docker Image Locally...")
	docker run -i -t --rm --name ${app} -p ${local_port}:3000 ${app}-golang:${version} 

test:
	$(call blue, "# Testing Golang Code...")
	docker run --rm -it -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${go_version} sh -c 'go test -v' 

clean: 
		@rm -f ${app} 

define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef

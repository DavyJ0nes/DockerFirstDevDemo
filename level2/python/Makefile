all: run_docker

image_name ?= docker-demo/python
local_port ?= 8082
py_version ?= 3.6.2
app_version ?= 0.0.1
git_hash = $(shell git rev-parse HEAD | cut -c 1-6)
build_date = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

build:
	$(call blue, "# Building Python App...")
	docker build --label APP_VERSION=${app_version} --label BUILT_ON=${build_date} --label GIT_HASH=${git_hash} -t ${image_name} .

run_docker: build
	$(call blue, "# Running Python Image...")
	docker run -p ${local_port}:5000 ${image_name}

test:
	$(call blue, "# Testing App Code...")
	docker run -v $(CURDIR):/srv/code -w /srv/code python:${py_version}-alpine3.6 pip3 install -r requirements.txt && python app_test.py

idle:
	$(call blue, "# Starting IDLE with v${py_version}...")
	docker run --rm -it python:${py_version}-alpine3.6


define blue
  @tput setaf 4
	@echo $1
	@tput sgr0
endef

all: run

image_name ?= docker-demo/python
local_port ?= 8082
py_version ?= 3.6.2

build:
	$(call blue, "# Building Python App...")
	docker build -t ${image_name} .

run_docker: build
	$(call blue, "# Running Python Image...")
	docker run -p ${local_port}:5000 ${image_name}

run:
	$(call blue, "# Running Code in Container...")
	docker run --rm -it python:${py_version}-alpine3.6

test:
	$(call blue, "# Testing App Code...")
	docker run -v $(CURDIR):/srv/code -w /srv/code python:${py_version}-alpine3.6 pip3 install -r requirements.txt && python app_test.py

idle:
	$(call blue, "# Starting IDLE with v${py_version}...")


define blue
  @tput setaf 4
	@echo $1
	@tput sgr0
endef

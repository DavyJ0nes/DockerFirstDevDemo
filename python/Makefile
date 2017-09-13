all: run

image_name ?= docker-demo/python
local_port ?= 8082

build:
	$(call blue, "# Building Python App...")
	docker build -t ${image_name} .

run: build
	$(call blue, "# Running Python Container...")
	docker run -p ${local_port}:5000 ${image_name}

define blue
  @tput setaf 4
	@echo $1
	@tput sgr0
endef

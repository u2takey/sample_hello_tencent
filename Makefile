
.PHONY: all 

PACKAGES = $(shell go list ./... | grep -v vendor/)
PWD = $(shell pwd)
IMAGE_NAME=ccr.ccs.tencentyun.com/kevinhub/sample-hello
BUILD_NUMBER=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shelldate +%FT%T%z)

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

build_local:
	mkdir -p make/release
	go build -o  make/release/hello github.com/u2takey/sample_hello_tencent

build_cross:
	GOOS=linux  GOARCH=amd64 CGO_ENABLED=0  go build -o  make/release/linux/amd64/hello github.com/u2takey/sample_hello_tencent

build_docker: build_cross
	docker build -t $(IMAGE_NAME):$(BUILD_NUMBER) .

push_docker: build_docker
	docker push $(IMAGE_NAME):$(BUILD_NUMBER)

run_local: build_local
	make/release/hello server

deploy_ccs:
	kubectl apply -f ./deploy.yaml
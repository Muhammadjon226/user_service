CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

ifneq (,$(wildcard ./.env))
	include .env
endif

ifdef CI_COMMIT_BRANCH
        include .build_info
endif

make create-env:
	cp ./.env.example ./.env

set-env:
	./scripts/set-env.sh ${CURRENT_DIR}

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh  ${CURRENT_DIR}

migrate-jeyran: set-env
	env POSTGRES_HOST=${POSTGRES_HOST} env POSTGRES_PORT=${POSTGRES_PORT} env POSTGRES_USER=${POSTGRES_USER} env POSTGRES_PASSWORD=${POSTGRES_PASSWORD} env POSTGRES_DB=${POSTGRES_DB} ./scripts/migrate-jeyran.sh

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --rebase

clear:
	rm -rf ${CURRENT_DIR}/bin/*

network:
	docker network create --driver=bridge ${NETWORK_NAME}

mark-as-staging-image:
	docker tag ${REGISTRY}/${IMG_NAME}:${TAG} ${REGISTRY}/${IMG_NAME}:staging
	docker push ${REGISTRY}/${IMG_NAME}:staging

mark-as-production-image:
	docker tag ${REGISTRY}/${IMG_NAME}:${TAG} ${REGISTRY}/${IMG_NAME}:production
	docker push ${REGISTRY}/${IMG_NAME}:production

build-image:
	docker build --rm -t ${REGISTRY}/${IMG_NAME}:${TAG} .

push-image:
	docker push ${REGISTRY}/${IMG_NAME}:${TAG}

run-dev: set-env
	go run cmd/main.go

dep:
	go get -v -d ./...

lint:
	golint -set_exit_status ${PKG_LIST}

unit-tests: set-env ## Run unit-tests
	go test -mod=vendor -v -cover -short ${PKG_LIST}

race: set-env ## Run data race detector
	go test -mod=vendor -race -short ${PKG_LIST}

msan: set-env ## Run memory sanitizer. If this test fails, you need to write the following command: export CC=clang (if you have installed clang)
	env CC=clang env CXX=clang++ go test -mod=vendor -msan -short ${PKG_LIST}

delete-branches:
	${CURRENT_DIR}/scripts/delete-branches.sh

run-localdb:
	docker run --name ${POSTGRES_DB} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -e POSTGRES_DB=${POSTGRES_DB} -e POSTGRES_USER=${POSTGRES_USER} -p ${POSTGRES_PORT}:5432 --restart always -d postgis/postgis 

.PHONY: vendor
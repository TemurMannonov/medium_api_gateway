CURRENT_DIR=$(shell pwd)

swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run cmd/main.go

local-up:
	docker compose --env-file ./.env.docker up -d

proto-gen:
	rm -rf genproto
	./scripts/gen-proto.sh ${CURRENT_DIR}

.PHONY: start
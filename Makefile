MAKEFLAGS += --silent

ifneq (,$(wildcard ./.env))
    include .env
    export
endif


export DB_URL= $(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

export DOCKER_CONNECTION= sudo docker exec ${container} bash -c

start:
	nodemon --exec go run main.go --signal SIGTERM 

prepare:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
	mv migrate /usr/bin &>/dev/null

migration:
	$(eval timestamp := $(shell date +%s))
	touch db/migrations/$(timestamp)_${name}.up.sql
	touch db/migrations/$(timestamp)_${name}.down.sql

dago:

SHELL := /bin/bash

dc := docker compose
de := $(dc) exec

.PHONY: config
build-config:
	@printf "Create application configs...\n"
	@cp ./.env.example ./.env
	@printf "Done!\n"

.PHONY: construct
build-construct: build-config
	$(dc) up -d --build

.PHONY: up
up:
	$(dc) up -d

.PHONY: down
down:
	$(dc) down

.PHONY: restart
restart:
	$(dc) down && $(dc) up

.PHONY: db-console
db-console:
	$(de) db $(SHELL)

.PHONY: app-console
app-console:
	$(de) app $(SHELL)

.PHONY: build a executive binary file
build-app:
	$(dc) up app --build

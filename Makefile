
# VARIABLES
AVAILABLE_DB_TYPES="sqlite, postgres, mysql"


# GLOBAL
export GO111MODULE=on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

run-mysql :		## Start MySQL container
	docker run -d --name mysql \
		-e MYSQL_ROOT_PASSWORD=supersecret \
		-p 3306:3306 \
		mysql

run-postgres :		## Start PostgreSQL container
	docker run -d --name postgres \
		-e POSTGRES_PASSWORD=supersecret \
		-p 5432:5432 \
		postgres

## applications

build :		## Build code base
	go build ./...

__check-db-type :
	@[ "$(DB_TYPE)" ] || ( echo "Missing database type (DB_TYPE), please define it and retry"; exit 1 )

run : __check-db-type		## Run application
	godotenv -f dotenvs/$(DB_TYPE).env go run main.go

## helpers

list-available-db-types :		## List all available database types
	@echo ""
	@echo "Available database types: $(AVAILABLE_DB_TYPES)"
	@echo ""

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""

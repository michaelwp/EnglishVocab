BUILD_DIR := ./bin
MIGRATION_PATH := migration/

.PHONY: build
build:
	@echo "build application is starting ..."
	@go build -o $(BUILD_DIR)
	@echo "... build application is finished"

.PHONY: run
run:
	@echo "application is started ..."
	@$(BUILD_DIR)/EnglishVocab

.phony: create-migrate-file
create-migrate-file:
	@migrate create -ext sql -dir migration/ -seq $(FILE_NAME)

.phony: migrate
migrate:
	@migrate -path $(MIGRATION_PATH) -database "postgresql://postgres:postgres@localhost:5432/database_name?sslmode=disable" -verbose up
APP_NAME := task
CMD_PATH := ./cmd/task/main.go

.PHONY: build run

build:
	go build -o $(APP_NAME) $(CMD_PATH)

run: build
	./$(APP_NAME)

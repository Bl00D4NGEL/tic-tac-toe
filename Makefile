.PHONY: build

build:
	go build -o bin/tic-tac-toe src/win-conditions.go src/main.go


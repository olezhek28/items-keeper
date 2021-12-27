.PHONY:
.SILENT:

build:
	go build -o ./bin/items_keeper_bot cmd/items_keeper_bot/main.go

run: build
	./bin/items_keeper_bot
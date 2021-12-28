.PHONY:
.SILENT:

build:
	go build -o ./bin/items_keeper_bot cmd/items_keeper_bot/main.go

run: build
	./bin/items_keeper_bot

build-image:
	docker build -t items_keeper_bot:v0.1 .

start-container:
	docker run --name items_keeper_bot -p 80:80 --env-file .env items_keeper_bot:v0.1
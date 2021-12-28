FROM golang:1.17.5-alpine3.15 AS builder

COPY . /github.com/olezhek28/items-keeper/
WORKDIR /github.com/olezhek28/items-keeper/

RUN go mod download
RUN go build -o ./bin/items_keeper_bot cmd/items_keeper_bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/olezhek28/items-keeper/bin/items_keeper_bot .
COPY --from=builder /github.com/olezhek28/items-keeper/configs/ configs/

EXPOSE 80

CMD ["./items_keeper_bot"]
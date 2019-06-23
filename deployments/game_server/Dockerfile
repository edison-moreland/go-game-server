# Expects context at repo root
FROM golang

ENV GO111MODULE=on

WORKDIR /app

RUN go get github.com/edison-moreland/go-game-server/game_server

EXPOSE 8080:8080
ENTRYPOINT ["game_server"]
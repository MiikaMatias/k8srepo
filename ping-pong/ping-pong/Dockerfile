FROM golang:latest

RUN mkdir /pinglogs
COPY ping-pong.go /
COPY go.mod go.sum /
RUN go mod download

CMD ["go", "run", "/ping-pong.go"]
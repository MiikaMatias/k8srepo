FROM golang:latest

RUN mkdir /logs
RUN mkdir /pinglogs
COPY reader.go /

CMD ["go", "run", "/reader.go"]
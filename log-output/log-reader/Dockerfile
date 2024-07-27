FROM golang:latest

RUN mkdir /logs /pinglogs /config
COPY reader.go /

CMD ["go", "run", "/reader.go"]
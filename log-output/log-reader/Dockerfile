FROM golang:latest

RUN mkdir /logs
COPY reader.go /

CMD ["go", "run", "/reader.go"]
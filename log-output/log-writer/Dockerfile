FROM golang:latest

RUN mkdir /logs
COPY writer.go /

CMD ["go", "run", "/writer.go"]